package checksec

import (
	"context"
	"xsploit/vul/cve-2099-9999"

	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/urfave/cli/v3"
)

const (
	CommandNameAuto = "auto"
)

var (
	Auto = &cli.Command{
		Name:    CommandNameAuto,
		Usage:   "auto",
		Aliases: []string{"a"},
		Action: func(ctx context.Context, cmd *cli.Command) (err error) {
			vulnerabilities := vul.Vulnerabilities{
				&cve_2099_9999.Vul,
			}
			err = vulnerabilities.Check(ctx)
			if err != nil {
				return
			}
			vulnerabilities.Output()
			return
		},
	}
)
