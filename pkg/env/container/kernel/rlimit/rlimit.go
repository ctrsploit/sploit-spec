package rlimit

import "golang.org/x/sys/unix"

type Resource struct {
	Name string `json:"name"`
	Type int
	Hard unix.Rlimit
	Soft unix.Rlimit
}

type Rlimits struct {
	Core       Resource `json:"core"`
	Cpu        Resource `json:"cpu"`
	Data       Resource `json:"data"`
	Fsize      Resource `json:"fsize"`
	Locks      Resource `json:"locks"`
	Msgqueue   Resource `json:"msgqueue"`
	Nice       Resource `json:"nice"`
	Rtprio     Resource `json:"rtprio"`
	Rttime     Resource `json:"rttime"`
	Sigpending Resource `json:"sigpending"`
	Stack      Resource `json:"stack"`
	As         Resource `json:"as"`
	Memlock    Resource `json:"memlock"`
	Nofile     Resource `json:"nofile"`
	Nproc      Resource `json:"nproc"`
	Rss        Resource `json:"rss"`
}
