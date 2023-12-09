package domain

type Configuration struct {
	Designs      bool `json:"designs"`
	Applications bool `json:"applications"`
	Filters      bool `json:"filters"`
}

// DesignerComponents represents a set of boolean flags for various components in a design or application context.
type DesignerComponents struct {
	Design      bool `json:"design,omitempty"`      // Indicates whether the design component is present.
	Application bool `json:"application,omitempty"` // Indicates whether the application component is present.
	Filter      bool `json:"filter,omitempty"`      // Indicates whether the filter component is present.
	Save        bool `json:"save,omitempty"`        // Indicates whether the save action is allowed.
	New         bool `json:"new,omitempty"`         // Indicates whether the new action is allowed.
	SaveAs      bool `json:"saveAs,omitempty"`      // Indicates whether the saveAs action is allowed.
	Validate    bool `json:"validate,omitempty"`    // Indicates whether the validate action is allowed.
	Deploy      bool `json:"deploy,omitempty"`      // Indicates whether the deploy action is allowed.
	Undeploy    bool `json:"undeploy,omitempty"`    // Indicates whether the undeploy action is allowed.
}
