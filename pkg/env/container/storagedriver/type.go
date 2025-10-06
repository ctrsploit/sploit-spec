package storagedriver

type Type int

const (
	TypeUnknown Type = iota
	TypeOverlay
	TypeDeviceMapper
	TypeAufs
)

func (t Type) String() string {
	switch t {
	case TypeOverlay:
		return "overlay"
	case TypeDeviceMapper:
		return "deviceMapper"
	case TypeAufs:
		return "aufs"
	default:
		return "unknown"
	}
}
