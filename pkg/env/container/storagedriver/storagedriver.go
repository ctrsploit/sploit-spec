package storagedriver

type StorageDriver struct {
	// driver type
	Type Type `json:"type"`
	// kernel module loaded
	Enabled bool `json:"enabled"`
	// used by container
	Used bool `json:"used"`
	// number of containers
	Number int `json:"number"`
	// host path of container rootfs
	Rootfs string `json:"rootfs"`
}
