package linux

type Env struct {
  TimeStamp string `json:"time_stamp"`
  ServiceName string `json:"service_name"`
	Basic
	LinuxSecurityFeature
	Advance
}