package linux

type LinuxSecurityFeature struct {
	Credential `json:"credential"`
	Unshare bool `json:"unshare"`
	Capability `json:"capability"`
	Seccomp   `json:"seccomp"`
	ProcFlag `json:"prcoflag"`
	LeakBase `json:"LeakBase"`
}

type Credential struct {
	Uid int `json:"uid"`
	Gid int `json:"gid"`
}

// type Unshare struct {
// 	Unshare bool `json:"unshare"`
// }


type Capability struct {
	Pid1 uint64 `json:"pid1"`
	Self uint64 `json:"self"`
}


type Seccomp struct {
	KernelSupported  bool   `json:"kernel_supported"`
	ContainerEnabled bool   `json:"container_enabled"`
	Mode             string `json:"mode"`
}

type ProcFlag struct {
	ProcFlag map[string]string `json:"proc_flag"`
}

type LeakBase struct {
	LeakBase string `json:"leak_base"`
}