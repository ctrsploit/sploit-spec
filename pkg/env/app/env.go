package app

type WebServerType int

const (
  Apache WebServerType = iota
  Nginx
  Tomcat
)
