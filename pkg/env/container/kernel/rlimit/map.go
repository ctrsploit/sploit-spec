package rlimit

import "golang.org/x/sys/unix"

var MapNameToType = map[string]int{
	"core":       unix.RLIMIT_CORE,       // max core file size
	"cpu":        unix.RLIMIT_CPU,        // CPU time in seconds
	"data":       unix.RLIMIT_DATA,       // max data size
	"fsize":      unix.RLIMIT_FSIZE,      // Maximum file size
	"locks":      unix.RLIMIT_LOCKS,      // max number of file locks
	"msgqueue":   unix.RLIMIT_MSGQUEUE,   // max bytes in POSIX message queues
	"nice":       unix.RLIMIT_NICE,       // max nice priority
	"rtprio":     unix.RLIMIT_RTPRIO,     // max real-time priority
	"rttime":     unix.RLIMIT_RTTIME,     // timeout for real-time tasks
	"sigpending": unix.RLIMIT_SIGPENDING, // max number of pending signals
	"stack":      unix.RLIMIT_STACK,      // max stack size
	"as":         unix.RLIMIT_AS,         // address space (virtual memory)
	"memlock":    unix.RLIMIT_MEMLOCK,    // max locked-in-memory address space
	"nofile":     unix.RLIMIT_NOFILE,     // max number of open files
	"nproc":      unix.RLIMIT_NPROC,      // max number of processes
	"rss":        unix.RLIMIT_RSS,        // max resident set size
}
