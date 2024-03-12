package app

type Env struct {
	WebServerType WebServerType `json:type`
	FrameworkType FrameworkType `json:type`
}
