package checksec

import (
	"xsploit/vul/cve-2099-9999"

	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:    "checksec",
	Aliases: []string{"c"},
	Usage:   "check security inside a container",
	Subcommands: []*cli.Command{
		Auto,
		app.Vul2ChecksecCmd(cve_2099_9999.Vul, []string{"2099"}, nil),
	},
}
