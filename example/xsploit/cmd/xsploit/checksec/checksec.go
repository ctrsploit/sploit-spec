package checksec

import (
	"xsploit/vul/cve-2099-9999"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:    "checksec",
	Aliases: []string{"c"},
	Usage:   "check security inside a container",
	Subcommands: []*cli.Command{
		Auto,
		cve_2099_9999.CheckSecCmd,
	},
}
