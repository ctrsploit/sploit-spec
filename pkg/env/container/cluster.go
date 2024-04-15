package container

type AccessType int

const (
	AccessDeny Type = iota
	AccessAnonymous
	AccessRead
	AccessWrite
	AccessExec
)

type CredentialType int

const (
	Certificate CredentialType = iota
	Token
)

type PodStatus int

const (
	Running PodStatus = iota
	Stop
)

type AccessRule struct {
	Resource string   `json:"resource"`
	Verbs    []string `json:"verbs"`
}

type K8sCredential struct {
	Type  CredentialType `json:"type"`
	User  string         `json:"user"`
	Rules []AccessRule   `json:"rules"`
}

type K8sPlugin struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type K8sPod struct {
	Name      string    `json:"name"`
	NameSpace string    `json:"namespace"`
	Status    PodStatus `json:"status"`
}

type K8sConfig struct {
	ConfigMap  map[string]map[string]string `json:config_map`
	Secret     map[string]map[string]string `json:secret`
	NameSpaces []string                     `json:namespaces`
	Pods       []K8sPod                     `json:pods`
}

type CustomResource struct {
	Name       string `json:"name"`
	ApiGroup   string `json:"api_group"`
	ApiVersion string `json:api_version`
}

type Cluster struct {
	Version         string            `json:"version"`
	ApiServer       string            `json:"apiserver"`
	PodEnv          map[string]string `json:"pod_env"`
	PodMount        map[string]string `json:"pod_mount"`
	DnsService      map[string]string `json:"dns_service"`
	ApiserverAccess AccessType        `json:"apiserver_access"`
	KubeletAccess   AccessType        `json:"kublet_access"`
	HostPath        map[string]string `json:"hostpath"`
	Credentials     []K8sCredential   `json:"credentials"`
	Plugins         []K8sPlugin       `json:"plugins"`
	Config          K8sConfig         `json:"config"`
	CustomResources []CustomResource  `json:"custom_resources"`
}
