package container

import (
	"github.com/moby/sys/mountinfo"
)

type Type struct {
	In    bool            `json:"in"`
	Rules map[string]bool `json:"rules"`
}

type Where struct {
	Container  Type `json:"container"`
	K8s        Type `json:"k8s"`
	Containerd Type `json:"containerd"`
	Docker     Type `json:"docker"`
	Nerdctl    Type `json:"nerdctl"`
}

type Basic struct {
	Where         Where             `json:"where"`
	KernelVersion string            `json:"kernel_version"`
	MountInfo     []*mountinfo.Info `json:"mountinfo"`
}
