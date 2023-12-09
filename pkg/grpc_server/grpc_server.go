package grpcserver

import (
	"context"
	"fmt"
	"net"

	"git.galaxymesh.io/pkg/tracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/nebula-aac/kubernexus-mono/api/protos/meshes"
	"github.com/nebula-aac/kubernexus-mono/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type NexusGRPCServer struct {
	meshes.UnimplementedNexusServiceServer
}

type Service struct {
	Name string `json:"name"`
	Port string `json:"port"`
	NexusGRPCServer
}

func grpcPanicHandler(r interface{}) error {
	fmt.Println("Panic occurred:", r)
	return errors.ErrPanic(r)
}

func Start(s *Service, _ tracing.Handler) error {
	address := fmt.Sprintf(":%s", s.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.ErrGrpcListener(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicHandler)),
		),
	)

	reflection.Register(server)

	meshes.RegisterNexusServiceServer(server, s)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	if err = server.Serve(listener); err != nil {
		return errors.ErrGrpcServer(err)
	}
	return nil
}
