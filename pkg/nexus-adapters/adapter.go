package adapters

import (
	"context"

	"github.com/nebula-aac/kubernexus-mono/api/protos/meshes"
	adapterconfig "github.com/nebula-aac/kubernexus-mono/pkg/adapter-config"
	"github.com/nebula-aac/kubernexus-mono/pkg/events"
	"github.com/nebula-aac/kubernexus-mono/pkg/nexuslogger"
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
	Config              adapterconfig.Handler
	KubeConfigHandler   adapterconfig.Handler
	Log                 nexuslogger.Logger
	CloudEventsStreamer *events.CloudEventsStreamer
}
