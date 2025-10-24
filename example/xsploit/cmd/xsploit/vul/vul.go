package vul

//goland:noinspection GoSnakeCaseUsage
import (
	cve_2099_9999 "xsploit/vul/cve-2099-9999"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:    "vul",
	Aliases: []string{"v"},
	Usage:   "list vulnerabilities",
	Subcommands: []*cli.Command{
		cve_2099_9999.VulCmd,
	},
}
