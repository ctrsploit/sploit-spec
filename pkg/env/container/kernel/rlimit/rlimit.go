package rlimit

import "golang.org/x/sys/unix"

type Rlimit struct {
	Core       unix.Rlimit `json:"core"`
	Cpu        unix.Rlimit `json:"cpu"`
	Data       unix.Rlimit `json:"data"`
	Fsize      unix.Rlimit `json:"fsize"`
	Locks      unix.Rlimit `json:"locks"`
	Msgqueue   unix.Rlimit `json:"msgqueue"`
	Nice       unix.Rlimit `json:"nice"`
	Rtprio     unix.Rlimit `json:"rtprio"`
	Rttime     unix.Rlimit `json:"rttime"`
	Sigpending unix.Rlimit `json:"sigpending"`
	Stack      unix.Rlimit `json:"stack"`
	As         unix.Rlimit `json:"as"`
	Memlock    unix.Rlimit `json:"memlock"`
	Nofile     unix.Rlimit `json:"nofile"`
	Nproc      unix.Rlimit `json:"nproc"`
	Rss        unix.Rlimit `json:"rss"`
}
