package domain

// Define the common provider domain entities and interfaces
type (
	// Represents provider types
	ProviderType string

	// Define the Provider's package version and endpoint
	ProviderPackage struct {
		PackageVersion  string `json:"package_version"`  // version of the package
		PackageEndpoint string `json:"package_endpoint"` // endpoint of the package
	}

	// Represents the Provider's structure and properties
	ProviderProperties struct {
		ProviderType        ProviderType     `json:"provider_type"`        // local or remote
		ProviderName        string           `json:"provider_name"`        // name of the provider
		ProviderDescription string           `json:"provider_description"` // description of the provider
		ProviderEndpoint    string           `json:"provider_endpoint"`    // endpoint of the provider
		ProviderPackage     ProviderPackage  `json:"provider_package"`     // package of the provider
		Extensions          Extensions       `json:"extensions"`           // extensions of the provider
		Capabilities        Capabilities     `json:"capabilities"`         // capabilities of the provider
		RestrictedAccess    RestrictedAccess `json:"restricted_access"`    // restricted access of the provider
	}

	ICommonProviders interface{}

	IProvider interface{}

	IRemoteProvider interface{}
)
