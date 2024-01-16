package checksec

import (
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/urfave/cli/v2"
	"xsploit/vul"
)

var Command = &cli.Command{
	Name:    "checksec",
	Aliases: []string{"c"},
	Usage:   "check security inside a container",
	Subcommands: []*cli.Command{
		Auto,
		app.Vul2ChecksecCmd(vul.CVE_2099_9999_v1, []string{"2099"}),
	},
}
