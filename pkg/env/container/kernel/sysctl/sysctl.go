package sysctl

type Sysctl struct {
	Net
	User
	Kernel
}

type Net struct {
	RouteLocalNet bool `json:"net.ipv4.conf.all.route_localnet"`
}

type User struct {
	MaxUserNamespaces int `json:"user.max_user_namespaces"`
}

type Kernel struct {
	UnprivilegedUsernsClone bool `json:"kernel.unprivileged_userns_clone"`
	PidMax                  int  `json:"kernel.pid_max"`
}
