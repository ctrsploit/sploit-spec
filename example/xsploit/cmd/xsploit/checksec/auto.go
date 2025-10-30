package checksec

import (
	"xsploit/vul/cve-2099-9999"

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
				&cve_2099_9999.Vul,
			}
			err = vulnerabilities.Check(context)
			if err != nil {
				return
			}
			vulnerabilities.Output()
			return
		},
	}
)
