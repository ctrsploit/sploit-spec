package app

import (
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/urfave/cli/v2"
)

func Vul2ChecksecCmd(v vul.Vulnerability, alias []string) *cli.Command {
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Action: func(ctx *cli.Context) (err error) {
			_, err = v.CheckSec(ctx)
			if err != nil {
				return
			}
			v.Output()
			return
		},
	}
}

func Vul2ExploitCmd(v vul.Vulnerability, alias []string) *cli.Command {
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Action: func(ctx *cli.Context) (err error) {
			_, err = v.Exploit(ctx)
			v.GetVulnerabilityExists()
			if err != nil {
				return
			}
			v.Output()
			return
		},
	}
}

func Vul2VulCmd(v vul.Vulnerability, alias []string) *cli.Command {
	checksec := Vul2ChecksecCmd(v, []string{"c"})
	checksec.Name = "checksec"
	checksec.Usage = "check vulnerability exists"

	exploit := Vul2ExploitCmd(v, []string{"x"})
	exploit.Name = "exploit"
	exploit.Usage = "run exploit"
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Subcommands: []*cli.Command{
			checksec,
			exploit,
		},
	}
}
