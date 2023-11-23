package checksec

import (
	vul2 "github.com/ctrsploit/sploit-spec/example/xsploit/vul"
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/urfave/cli/v2"
)

const (
	CommandNameAuto = "auto"
)

var (
	Auto = &cli.Command{
		Name:    CommandNameAuto,
		Usage:   "auto",
		Aliases: []string{"a"},
		Action: func(context *cli.Context) (err error) {
			vulnerabilities := vul.Vulnerabilities{
				vul2.CVE_2099_9999_v1,
			}
			err = vulnerabilities.Check()
			if err != nil {
				return
			}
			vulnerabilities.Output()
			return
		},
	}
)
