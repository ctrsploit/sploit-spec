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
	hyper_type          string `json:"hyper_type"`
	qemu_ver 			string `json:"qemu_ver"`
	dev_list 	     [] DeviceInfo `json:"dev_list"`
}

type DeviceInfo struct {
	type 	DeviceType 	`json:"type"`
	name 	string	 	`json:"name"`
}
