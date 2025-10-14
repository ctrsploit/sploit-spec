package container

type Kernel struct {
	Sysctl `json:"sysctl"`
}

type Sysctl struct {
	ProcSysNetIpv4ConfAllRouteLocalNet bool `json:"route_localnet"`
}
