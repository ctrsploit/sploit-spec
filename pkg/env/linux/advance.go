package linux

type Advance struct {
	EnvVars []string `json:"env_vars"`
	SuidBinarys []string `json:"sudi_binarys"`
	SudoVersion string `json:"sudo_version"`
}