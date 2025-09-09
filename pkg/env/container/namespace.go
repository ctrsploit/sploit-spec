package container

import "github.com/ctrsploit/sploit-spec/internal"

// Name

const (
	NamespaceNameCGroup          = "cgroup"
	NamespaceNameIpc             = "ipc"
	NamespaceNameMnt             = "mnt"
	NamespaceNameNet             = "net"
	NamespaceNamePid             = "pid"
	NamespaceNamePidForChildren  = "pid_for_children"
	NamespaceNameUser            = "user"
	NamespaceNameUts             = "uts"
	NamespaceNameTime            = "time"
	NamespaceNameTimeForChildren = "time_for_children"
)

type NamespaceLevel int

const (
	NamespaceLevelUnknown NamespaceLevel = iota
	NamespaceLevelBoot
	NamespaceLevelChild
	NamespaceLevelNotSupported
	NamespaceLevelHost = NamespaceLevelBoot
)

var (
	NamespaceLevelMap = map[NamespaceLevel]string{
		NamespaceLevelChild:        "child",
		NamespaceLevelHost:         "host",
		NamespaceLevelNotSupported: "not supported( <=> host)",
		NamespaceLevelUnknown:      "unknown",
	}
)

func (l NamespaceLevel) String() string {
	return NamespaceLevelMap[l]
}

// Type

type NamespaceType int

const (
	NamespaceTypeUnknown NamespaceType = iota
	NamespaceTypeIPC
	NamespaceTypeUTS
	NamespaceTypeUser
	NamespaceTypePid
	NamespaceTypeCGroup
	NamespaceTypeTime
	NamespaceTypeMount
	NamespaceTypeNetwork
)

// Map

var (
	NamespaceMapName2Type = map[string]NamespaceType{
		NamespaceNameCGroup: NamespaceTypeCGroup,
		NamespaceNameIpc:    NamespaceTypeIPC,
		NamespaceNameMnt:    NamespaceTypeMount,
		NamespaceNameNet:    NamespaceTypeNetwork,
		NamespaceNamePid:    NamespaceTypePid,
		// TODO: not sure pid_for_children is same as pid?
		NamespaceNamePidForChildren: NamespaceTypePid,
		NamespaceNameUser:           NamespaceTypeUser,
		NamespaceNameUts:            NamespaceTypeUTS,
		NamespaceNameTime:           NamespaceTypeTime,
		// TODO: not sure time_for_children is same as time?
		NamespaceNameTimeForChildren: NamespaceTypeTime,
	}
	NamespaceMapType2Name = internal.ReverseMap(NamespaceMapName2Type).(map[NamespaceType]string)
)

func (l NamespaceType) String() string {
	return NamespaceMapType2Name[l]
}
