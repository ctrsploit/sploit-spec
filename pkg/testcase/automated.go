package testcase

import "github.com/ctrsploit/sploit-spec/pkg/prerequisite"

type Automated struct {
	// Usage describes how to reproduce the same effects as the manual
	// test steps using automation commands or scripts. It usually includes
	// command-line examples, scripts, or tool configurations that can
	// perform the test automatically.
	Usage              string `json:"usage"`
	CommandName        string
	CommandDescription string
	CheckSec           prerequisite.BasePrerequisite
}
