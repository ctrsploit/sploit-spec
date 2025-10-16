package vul

import (
	"xsploit/vul/cve-2099-9999"

	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:    "vul",
	Aliases: []string{"v"},
	Usage:   "list vulnerabilities",
	Subcommands: []*cli.Command{
		app.Vul2VulCmd(cve_2099_9999.Vul, []string{"2099"}, nil, nil, true),
	},
}
