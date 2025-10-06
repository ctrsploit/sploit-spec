package container

type LinuxSecurityFeature struct {
	Credential   `json:"credential"`
	Capabilities `json:"capability"`
	LSM
	Seccomp   `json:"seccomp"`
	Namespace `json:"namespace"`
	CGroups   `json:"cgroups"`
	Filesystem
}

type Credential struct {
	Uid int `json:"uid"`
	Gid int `json:"gid"`
}

type Capability struct {
	Eff uint64 `json:"eff"`
	Bnd uint64 `json:"bnd"`
}

type Capabilities struct {
	Pid1 Capability `json:"pid1"`
	Self Capability `json:"self"`
}

type LSM struct {
	Apparmor `json:"apparmor"`
	SELinux  `json:"selinux"`
}

type Apparmor struct {
	KernelSupported  bool   `json:"kernel_supported"`
	ContainerEnabled bool   `json:"container_enabled"`
	Profile          string `json:"profile"`
	Mode             string `json:"mode"`
}

type SELinux struct {
	KernelSupported  bool   `json:"kernel_supported"`
	ContainerEnabled bool   `json:"container_enabled"`
	Mode             string `json:"mode"`
	MountPoint       string `json:"mount_point"`
}

type Seccomp struct {
	KernelSupported  bool   `json:"kernel_supported"`
	ContainerEnabled bool   `json:"container_enabled"`
	Mode             string `json:"mode"`
}

type Namespace struct {
	Levels map[string]NamespaceLevel `json:"levels"`
	Names  []string                  `json:"-"`
}

type CgroupsVersion int

const (
	CGroupsUnknown = iota
	CgroupsV1
	CgroupsV2
)

type CGroups struct {
	Version            CgroupsVersion `json:"version"`
	Subsystems         []string       `json:"sub"`
	TopLevelSubSystems []string       `json:"top"`
}

type Filesystem struct {
	StorageDriver string `json:"storage_driver"`
}
