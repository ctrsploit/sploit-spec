package container

type Type int

const (
	NotInContainer Type = iota
	K8s
	Containerd
	Docker
)

type Basic struct {
	Type          Type   `json:"type"`
	KernelVersion string `json:"kernel_version"`
}
