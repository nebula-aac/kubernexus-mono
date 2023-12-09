package domain

type (
	ContextKey string
	Feature    string
)

const (
	LocalProviderType  ProviderType = "local"
	RemoteProviderType ProviderType = "remote"
)

const (
	ProviderCtxKey       ContextKey = "provider"
	TokenCtxKey          ContextKey = "token"
	UserCtxKey           ContextKey = "user"
	PerformanceObjCtxKey ContextKey = "performance_obj"
	KubeClustersKey      ContextKey = "kube_clusters"
	AllKubeClusterKey    ContextKey = "all_kube_clusters"
)

const (
	MesheryServerUrlKey         ContextKey = "meshery_server_url"
	MesheryServerCallbackUrlKey ContextKey = "meshery_server_callback_url"
)

const (
	UserIdentity    Feature = "user_identity"
	UserKey         Feature = "user_key"
	UserProfile     Feature = "user_profile"
	SyncPreferences Feature = "sync_preferences"
	PersistResults  Feature = "persist_results"
	PersistResult   Feature = "persist_result"
)
