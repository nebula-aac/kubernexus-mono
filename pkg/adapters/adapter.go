package adapters

import (
	"context"

	"github.com/nebula-aac/kubernexus-mono/api/protos/meshes"
	"github.com/nebula-aac/kubernexus-mono/pkg/config"
	"github.com/nebula-aac/kubernexus-mono/pkg/events"
	"github.com/nebula-aac/kubernexus-mono/pkg/logger"
)

type Handler interface {
	GetName() string
	GetComponentInfo(interface{}) error
	ApplyOperation(context.Context, OperationRequest) error
	ListOperations() (Operations, error)
	ProcessOAM(ctx context.Context) (string, error)
	StreamErr(*meshes.CloudEventsResponse, error)
	StreamInfo(*meshes.CloudEventsResponse)
}

type Adapter struct {
	Config              config.Handler
	KubeConfigHandler   config.Handler
	Log                 logger.Logger
	CloudEventsStreamer *events.CloudEventsStreamer
}
