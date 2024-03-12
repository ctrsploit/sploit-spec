package app

type Basic struct {
	WebServer WebServer `json:"web_server"`
	Framework Framework `json:"framework"`
}

type WebServer struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Framework struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
