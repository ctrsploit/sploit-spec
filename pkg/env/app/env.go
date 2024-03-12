package app

type WebServerType int
type FrameworkType int

const (
	WebServerForUnknown WebServerType = iota
	WebServerForOther
	WebServerForApache
	WebServerForNginx
	WebServerForTomcat
	WebServerForIIS
)

const (
	FrameworkForUnknown FrameworkType = iota
	FrameworkForOther
	FrameworkForJavaSpring
	FrameworkForPyFlask
	FrameworkForPyDjango
	FrameworkForGoBeego
	FrameworkForGoGin
)
