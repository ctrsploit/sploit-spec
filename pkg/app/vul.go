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
		Action: func(context *cli.Context) (err error) {
			_, err = v.CheckSec()
			if err != nil {
				return
			}
			v.Output()
			return
		},
	}
}

func Vul2ExploitCmd(v vul.Vulnerability, alias []string, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Flags:   flags,
		Action: func(context *cli.Context) (err error) {
			_, err = v.CheckSec()
			if err != nil {
				return
			}
			err = v.Exploit()
			return
		},
	}
}

func Vul2VulCmd(v vul.Vulnerability, alias []string, flags []cli.Flag) *cli.Command {
	checksec := Vul2ChecksecCmd(v, []string{"c"})
	checksec.Name = "checksec"
	checksec.Usage = "check vulnerability exists"

	exploit := Vul2ExploitCmd(v, []string{"x"}, flags)
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
