package risk

type Risk int

// sort by risk level, from low to high
const (
	HostInformationDisclosure = 1 << iota
	LocalPrivilegeEscalate
	ContainerEscape
	UndefinedRisk Risk = 0
)
