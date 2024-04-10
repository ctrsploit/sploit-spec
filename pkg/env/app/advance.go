package app

type Advance struct {
	OS            OS          `json:"os"`
	ComponentList []Component `json:"component_list"`
}

type OS struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Component struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
