package kernel

import (
	"time"

	"github.com/ctrsploit/sploit-spec/pkg/env/container/kernel/rlimit"
	"github.com/ctrsploit/sploit-spec/pkg/env/container/kernel/sysctl"
)

type Kernel struct {
	Basic
	sysctl.Sysctl `json:"sysctl"`
	rlimit.Rlimit `json:"rlimit"`
}

type Basic struct {
	CompiledDate  time.Time `json:"compiled_date"`
	KernelVersion string    `json:"kernel_version"`
}
