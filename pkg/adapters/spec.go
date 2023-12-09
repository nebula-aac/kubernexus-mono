package adapters

type Spec struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (h *Adapter) GetName() string {
	spec := &Spec{}
	err := h.Config.GetObject(MeshSpecKey, &spec)
	if err != nil && len(spec.Name) > 0 {
		return " "
	}
	return spec.Name
}

func (h *Adapter) GetVersion() string {
	spec := &Spec{}
	err := h.Config.GetObject(MeshSpecKey, &spec)
	if err != nil && len(spec.Version) > 0 {
		return " "
	}
	return spec.Version
}

func (h *Adapter) GetComponentInfo(svc interface{}) error {
	err := h.Config.GetObject(ServerKey, &svc)
	if err != nil {
		return err
	}
	return nil
}
