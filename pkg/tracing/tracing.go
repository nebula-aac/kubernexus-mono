package tracing

import (
	"context"
	"errors"
	"reflect"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type KeyValue struct {
	Key   string
	Value string
}

type Event interface{}

// Handler represents a tracing handler that provides functionality
// to create tracers, spans and add events.
type Handler interface {
	// Tracer creates a new tracer with the given name.
	Tracer(ctx context.Context, name string) trace.Tracer

	// Span creates a new span with the context.
	Span(ctx context.Context, handler func(context.Context))

	// AddEvent adds a new event with the given name and attributes to the span.
	AddEvent(ctx context.Context, event Event)
}

type handler struct {
	provider trace.TracerProvider
	context  context.Context
	span     trace.Span
	tracer   trace.Tracer
	mu       sync.Mutex
}

func New(ctx context.Context, service string, endpoint string) (*handler, error) {
	if len(endpoint) < 2 {
		return nil, errors.New("invalid endpoint provided")
	}

	provider, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(provider)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(provider),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
		)),
	)

	tracer := tp.Tracer(service)

	return &handler{
		provider: tp,
		tracer:   tracer,
	}, err
}

func (h *handler) Tracer(ctx context.Context, name string) trace.Tracer {

	// Get span context
	spanCtx := trace.SpanContextFromContext(ctx)

	// Create tracer
	tracer := h.provider.Tracer(name)

	// Update the existing ctx to contain the span context
	ctx = trace.ContextWithSpanContext(ctx, spanCtx)

	// Return only the tracer
	return tracer
}

func (h *handler) Span(ctx context.Context, handler func(ctx context.Context)) {

	span := trace.SpanFromContext(ctx)
	defer span.End()

	handler(ctx)

}

func (h *handler) AddEvent(ctx context.Context, event Event) {

	// Lock mutex
	h.mu.Lock()
	defer h.mu.Unlock()

	// Build event options
	options := make([]trace.EventOption, 0)

	// Use reflection to get attributes
	val := reflect.ValueOf(event)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			options = append(options, trace.WithAttributes(
				attribute.String(field.Type().Name(), field.String())))
		}
	}

	// Use reflection to get event name
	name := reflect.TypeOf(event).Name()

	// Add event
	h.span.AddEvent(name, options...)

}
