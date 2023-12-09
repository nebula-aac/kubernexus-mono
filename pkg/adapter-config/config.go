package adapterconfig

type Handler interface {
	SetKey(key string, value string)
	GetKey(key string) string
	GetObject(key string, result interface{}) error
	SetObject(key string, value interface{}) error
}
