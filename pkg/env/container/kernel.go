package container

type Kernel struct {
	Sysctl `json:"sysctl"`
}

type Sysctl struct {
	Net
	User
}

type Net struct {
	RouteLocalNet bool `json:"net.ipv4.conf.all.route_localnet"`
}

type User struct {
	MaxUserNamespaces int `json:"user.max_user_namespaces"`
}
