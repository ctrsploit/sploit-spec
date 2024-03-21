package app

type Advance struct {
	OSType        string      `json:"os_type"`
	ComponentList []Component `json:"component_list"`
}

type Component struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
