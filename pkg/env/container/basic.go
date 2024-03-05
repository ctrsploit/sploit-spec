package container

type ContainerType int

const (
  NotInContainer ContainerType = iota
  K8s
  Containerd
  Docker
)
