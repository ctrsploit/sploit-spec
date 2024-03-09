package vul

import (
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/urfave/cli/v2"
	"xsploit/vul"
)

var Command = &cli.Command{
	Name:    "vul",
	Aliases: []string{"v"},
	Usage:   "list vulnerabilities",
	Subcommands: []*cli.Command{
		app.Vul2VulCmd(vul.CVE_2099_9999_v1, []string{"2099"}),
	},
}
