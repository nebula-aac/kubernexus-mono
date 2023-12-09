package domain

type (
	// Indicate to Provider whether there is a feature or not
	Capability struct {
		Feature  string `json:"feature"`
		Endpoint string `json:"endpoint"`
	}

	Capabilities []Capability
)

type HeaderComponents struct{}

type NavigatorComponents struct{}

type UserInterfaceCapabilities struct{}

type RestrictedAccess struct {
	AllowedComponents UserInterfaceCapabilities `json:"allowed_components"`
}
