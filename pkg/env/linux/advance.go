package linux

type Advance struct {
	EnvVars []string `json:"env_vars"`
	SuidBinarys []string `json:"sudi_binarys"`
	BinaryVersions map[string]string `json:"binary_versions"`
}