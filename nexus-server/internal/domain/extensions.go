package domain

import "net/http"

// Define the common extensions domain entities and interfaces
type (
	Router struct {
		HTTPHandler http.Handler
		Path        string
	}

	// Input for a plugin
	ExtensionInput struct {
		Router *Router
	}

	// Output for a plugin
	ExtensionOutput struct{}

	CommonComponent struct {
		Component string `json:"component"`
		Type      string `json:"type"`
	}

	CommonUserInterfaceExtension struct {
		Title           string `json:"title"`
		Href            string `json:"href"`
		Show            *bool  `json:"show"`
		Type            string `json:"type"`
		OnClickCallback int    `json:"onClickCallback"`
		CommonComponent
	}

	NavigatorExtension struct {
		Icon     string              `json:"icon"`
		Link     *bool               `json:"link"`
		IsBeta   *bool               `json:"isBeta"`
		Children NavigatorExtensions `json:"children"`
		CommonUserInterfaceExtension
	}

	UserPreferencesExtension struct {
		CommonComponent
	}
	GraphQLExtension struct {
		Path string `json:"path"`
		CommonComponent
	}
	AccountExtension struct {
		Children AccountExtensions `json:"children"`
		CommonUserInterfaceExtension
	}
	CollaboratorExtension struct {
		CommonComponent
	}

	NavigatorExtensions       []NavigatorExtension
	UserPreferencesExtensions []UserPreferencesExtension
	GraphQLExtensions         []GraphQLExtension
	AccountExtensions         []AccountExtension
	CollaboratorExtensions    []CollaboratorExtension

	// Defines the user interface entry points
	Extensions struct {
		Navigator       NavigatorExtensions       `json:"navigator"`
		UserPreferences UserPreferencesExtensions `json:"userPreferences"`
		GraphQl         GraphQLExtensions         `json:"graphQl"`
		Account         AccountExtensions         `json:"account"`
		Collaborator    CollaboratorExtensions    `json:"collaborator"`
	}
)
