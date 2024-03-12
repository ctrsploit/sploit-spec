package container

type Advance struct {
	RuntimeVersion `json:"runtime_version"`
	CtrCnt         `json:"ctr_cnt"`
}

type RuntimeVersion struct {
}

type CtrCnt struct {
}
