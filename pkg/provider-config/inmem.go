package providerconfig

import (
	adapterconfig "github.com/nebula-aac/kubernexus-mono/pkg/adapter-config"
	adapters "github.com/nebula-aac/kubernexus-mono/pkg/nexus-adapters"
	"github.com/nebula-aac/kubernexus-mono/pkg/utils"
)

type InMem struct {
	store        map[string]string
	marshaller   utils.Marshaller
	unmarshaller utils.Unmarshaller
}

// NewAdapterInMem returns a new instance of an in-memory configuration provider using the provided Options opts.
func NewAdapterInMem(opts InMemOptions) (adapterconfig.Handler, error) {
	store := make(map[string]string)

	val, err := opts.Marshaller.Marshal(opts.ServerConfig)
	if err != nil {
		return nil, err
	}
	store[adapters.ServerKey] = val

	val, err = opts.Marshaller.Marshal(opts.MeshSpec)
	if err != nil {
		return nil, err
	}
	store[adapters.MeshSpecKey] = val

	val, err = opts.Marshaller.Marshal(opts.Operations)
	if err != nil {
		return nil, err
	}
	store[adapters.OperationsKey] = val

	return &InMem{
		store: store,
	}, nil
}

// ---------------------------------Application config methods----------------------------------

// SetKey sets a key value in local store.
func (l *InMem) SetKey(key string, value string) {
	l.store[key] = value
}

// GetKey gets a key value from local store.
func (l *InMem) GetKey(key string) string {
	return l.store[key]
}

// GetObject gets an object value for the key.
func (l *InMem) GetObject(key string, result interface{}) error {
	return l.unmarshaller.Unmarshal(l.store[key], result)
}

// SetObject sets an object value for the key.
func (l *InMem) SetObject(key string, value interface{}) error {
	val, err := l.marshaller.Marshal(value)
	if err != nil {
		return err
	}
	l.store[key] = val
	return nil
}
