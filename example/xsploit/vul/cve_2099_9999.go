package vul

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite/user"
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	prerequisite2 "xsploit/pkg/prerequisite"
)

type CVE_2099_9999 struct {
	vul.BaseVulnerability
}

var (
	CVE_2099_9999_v1 = &CVE_2099_9999{
		vul.BaseVulnerability{
			Name:        "CVE-2099-9999",
			Description: "Description of CVE-2099-9999",
			CheckSecPrerequisites: prerequisite.Prerequisites{
				&prerequisite2.EvenTime,
			},
			ExploitablePrerequisites: prerequisite.Prerequisites{
				&user.MustBeRoot,
			},
		},
	}
)

func (cve CVE_2099_9999) Exploit() (err error) {
	err = cve.BaseVulnerability.Exploit()
	if err != nil {
		return
	}
	fmt.Println("CVE-2099-9999 has exploited")
	return
}
