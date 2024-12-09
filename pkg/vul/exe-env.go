package vul

type ExeEnv int

const (
	Unknown ExeEnv = iota
	Local
	Remote
	InContainer // execute in the container
	K8S         // execute with kubeconfig
)
