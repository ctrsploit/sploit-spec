package vul

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite/vulnerability"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/urfave/cli/v2"
)

type Vulnerability interface {
	// GetName returns a one word name; may be used as command name
	GetName() string
	// GetDescription return usage
	GetDescription() string
	GetVulnerabilityExists() bool
	GetVulnerabilityResponse() string
	Info()
	// CheckSec whether vulnerability exists
	CheckSec(ctx *cli.Context) (bool, error)
	// Output shows checksec result
	Output()
	// Exploitable whether vulnerability can be exploited,
	// will be called automatically before Exploit()
	Exploitable() (bool, error)
	Exploit(ctx *cli.Context) (bool, error)
}

type BaseVulnerability struct {
	Name                     string                     `json:"name"`
	Description              string                     `json:"description"`
	VulnerabilityExists      bool                       `json:"vulnerability_exists"`
	VulnerabilityResponse    string                     `json:"vulnerability_response"`
	CheckSecHaveRan          bool                       `json:"-"`
	CheckSecPrerequisites    prerequisite.Prerequisites `json:"-"`
	ExploitablePrerequisites prerequisite.Prerequisites `json:"-"`
}

func (v *BaseVulnerability) GetName() string {
	return v.Name
}

func (v *BaseVulnerability) GetDescription() string {
	return v.Description
}

func (v *BaseVulnerability) GetVulnerabilityExists() bool {
	return v.VulnerabilityExists
}

func (v *BaseVulnerability) GetVulnerabilityResponse() string {
	return v.VulnerabilityResponse
}

func (v *BaseVulnerability) Info() {
	log.Logger.Info(v.Description)
}

func (v *BaseVulnerability) CheckSec(ctx *cli.Context) (vulnerabilityExists bool, err error) {
	vulnerabilityExists, err = v.CheckSecPrerequisites.Satisfied()
	if err != nil {
		return
	}
	v.VulnerabilityExists = vulnerabilityExists
	v.CheckSecHaveRan = true
	return
}

func (v *BaseVulnerability) Output() {
	result := item.Resp{
		Name:        v.GetName(),
		Description: v.GetDescription(),
		Result:      v.GetVulnerabilityExists(),
		Response:    v.GetVulnerabilityResponse(),
	}
	fmt.Println(printer.Printer.Print(result))
}

func (v *BaseVulnerability) Exploitable() (satisfied bool, err error) {
	if !v.CheckSecHaveRan {
		panic(fmt.Errorf("CheckSecHaveRan = %+v", v.CheckSecHaveRan))
	}
	prerequisiteVulnerabilityExists := vulnerability.Exists(v.VulnerabilityExists)
	v.ExploitablePrerequisites = append([]prerequisite.Interface{prerequisiteVulnerabilityExists}, v.ExploitablePrerequisites...)
	satisfied, err = v.ExploitablePrerequisites.Satisfied()
	if err != nil {
		return
	}
	return
}

func (v *BaseVulnerability) Exploit(ctx *cli.Context) (vulnerabilityExists bool, err error) {
	exploitable, err := v.Exploitable()
	if err != nil {
		return
	}
	if !exploitable {
		err = fmt.Errorf("%s is not exploitable", v.Name)
		awesome_error.CheckErr(err)
		return
	}
	// Implemented it
	return
}
