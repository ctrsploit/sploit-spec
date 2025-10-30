package container

import "time"

type Kernel struct {
	KernelBasic
	Sysctl `json:"sysctl"`
}

type KernelBasic struct {
	CompiledDate  time.Time `json:"compiled_date"`
	KernelVersion string    `json:"kernel_version"`
}

type Sysctl struct {
	Net
	User
	KernelSysctl
}

type Net struct {
	RouteLocalNet bool `json:"net.ipv4.conf.all.route_localnet"`
}

type User struct {
	MaxUserNamespaces int `json:"user.max_user_namespaces"`
}

type KernelSysctl struct {
	UnprivilegedUsernsClone bool `json:"kernel.unprivileged_userns_clone"`
	PidMax                  int  `json:"kernel.pid_max"`
}
