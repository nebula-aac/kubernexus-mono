package configprovider

import (
	"github.com/nebula-aac/kubernexus-mono/pkg/adapters"
	"github.com/nebula-aac/kubernexus-mono/pkg/utils"
)

const (
	ViperKey = "viper"
	InMemKey = "in-mem"
)

/*
type Options struct {
	ServerConfig   map[string]string   // ServerConfig options are used configure the gRPC service of the adapter.
	MeshSpec       map[string]string   // MeshSpec options are used to configure the service mesh to be used.
	ProviderConfig map[string]string   // ProviderConfig options are used to configure the config provider.
	Operations     adapters.Operations // Operations contains the properties of the operations the adapter supports.
}
*/

type ViperOptions struct {
	FilePath string
	FileType string
	FileName string
}

type InMemOptions struct {
	Marshaller   utils.Marshaller
	Unmarshaller utils.Unmarshaller
	Options
}

type Options struct {
	Viper        *ViperOptions
	InMem        *InMemOptions
	ServerConfig map[string]string
	MeshSpec     map[string]string
	Operations   map[string]*adapters.Operation
}
