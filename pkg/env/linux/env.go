package kernel

type Env struct {
  Type ContainerType `json:type`
}

type SystemInfo struct {
	Times string `json:"Times"`
  // Type KernelType `json:type`
  Service_Name string `json:"Service_Name"`    
  Kernel_Version string `json:"Kernel_Version"`
  Compile_Time string `json:"Compile_Time"`
  UID string `json:"UID"`
  Unshare string `json:"Unshare"`
  Enable_Seccomp string `json:"Enable_Seccomp"`
  Support_Seccomp string `json:"Support_Seccomp"`
  Current_Process_Caps string `json:"Current_Process_Caps"`
  Pid1_Caps string `json:"Pid1_Caps"`
  UnPrivileged_bpf_disable bool `"json:UnPrivileged_bpf_disable"`
  Kernel_Base string `"json:Kernel_Base"`
  Environ string `"json:Environ"`
  Suid_List []string `"json:Suid_List"`
    //Class *Class `json:"class"`
}