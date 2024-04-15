package container

type ContainerType int

const (
	NotInContainer ContainerType = iota
	K8s
	Containerd
	Docker
)

type Basic struct {
	Type          ContainerType `json:"type"`
	KernelVersion string        `json:"kernel_version"`
}
