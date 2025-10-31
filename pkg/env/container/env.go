package container

import "github.com/ctrsploit/sploit-spec/pkg/env/container/kernel"

type Env struct {
	Basic
	kernel.Kernel
	LinuxSecurityFeature
	Cluster
	Advance
}
