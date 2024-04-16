package container

type Type struct {
	In    bool            `json:"in"`
	Rules map[string]bool `json:"rules"`
}

type Where struct {
	Container  Type `json:"container"`
	K8s        Type `json:"k8s"`
	Containerd Type `json:"containerd"`
	Docker     Type `json:"docker"`
}

type Basic struct {
	Where         Where  `json:"where"`
	KernelVersion string `json:"kernel_version"`
}
