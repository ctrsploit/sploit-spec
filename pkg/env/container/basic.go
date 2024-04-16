package container

type Where struct {
	NotInContainer bool
	K8s            bool
	Containerd     bool
	Docker         bool
}

type Basic struct {
	Where         Where  `json:"where"`
	KernelVersion string `json:"kernel_version"`
}
