package events

import (
	"context"
	"sync"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudEventsStreamer struct {
	clientChannels []chan cloudevents.Event
	clmx           sync.RWMutex
	cancelCtx      context.Context
	cancelFunc     context.CancelFunc
}

func NewCloudEventsStreamer() *CloudEventsStreamer {
	ctx, cancel := context.WithCancel(context.Background())
	return &CloudEventsStreamer{
		clientChannels: make([]chan cloudevents.Event, 0),
		cancelCtx:      ctx,
		cancelFunc:     cancel,
	}
}

func (e *CloudEventsStreamer) Subscribe(ch chan cloudevents.Event) {
	e.clmx.Lock()
	defer e.clmx.Unlock()
	e.clientChannels = append(e.clientChannels, ch)
}

func (e *CloudEventsStreamer) Unsubscribe(ch chan cloudevents.Event) {
	e.clmx.Lock()
	defer e.clmx.Unlock()
	for i, c := range e.clientChannels {
		if c == ch {
			e.clientChannels = append(e.clientChannels[:i], e.clientChannels[i+1:]...)
			return
		}
	}
}

func (e *CloudEventsStreamer) Publish(event cloudevents.Event) {
	e.clmx.RLock()
	defer e.clmx.RUnlock()
	for _, ch := range e.clientChannels {
		ch <- event
	}
}

func (e *CloudEventsStreamer) Stop() {
	e.cancelFunc()
}
