package exeenv

const (
	Local       = 1 << 0
	Remote      = 1 << 1
	InContainer = 1 << 2 // execute in the container
	K8S         = 1 << 3 // execute with kubeconfig
)
