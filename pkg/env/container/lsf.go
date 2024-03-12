package container

type LinuxSecurityFeature struct {
	Credential `json:"credential"`
	Capability `json:"capability"`
	LSM
	Seccomp    `json:"seccomp"`
	Namespace  `json:"namespace"`
	CGroups    `json:"cgroups"`
	Filesystem `json:"filesystem"`
}

type Credential struct {
	Uid int `json:"uid"`
	Gid int `json:"gid"`
}

type Capability struct {
	Pid1 uint64 `json:"pid1"`
	Self uint64 `json:"self"`
}

type LSM struct {
	Apparmor `json:"apparmor"`
	SELinux  `json:"selinux"`
}

type Apparmor struct {
}

type SELinux struct {
}

type Seccomp struct {
}

type Namespace struct {
}

type CGroups struct {
}

type Filesystem struct {
}
