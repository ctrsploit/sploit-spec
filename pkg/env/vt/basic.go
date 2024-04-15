package vt


type DeviceType int

const (
	Disk DeviceType = iota
	Network
	Display
	USB
	Memory
	Other
)  

type Basic struct {
	HyperType          string `json:"hyper_type"`
	QemuVer 	   string `json:"qemu_ver"`
	DevList 	   [] DeviceInfo `json:"dev_list"`
}

type DeviceInfo struct {
	Type 	DeviceType 	`json:"type"`
	Name 	string	 	`json:"name"`
}
