package app

import (
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/urfave/cli/v2"
)

func Vul2ChecksecCmd(v vul.Vulnerability, alias []string, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Flags:   flags,
		Action: func(context *cli.Context) (err error) {
			_, err = v.CheckSec(context)
			if err != nil {
				return
			}
			v.Output()
			return
		},
	}
}

func Vul2ExploitCmd(v vul.Vulnerability, alias []string, flags []cli.Flag, checkBeforeExploit bool) *cli.Command {
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Flags:   flags,
		Action: func(context *cli.Context) (err error) {
			if checkBeforeExploit {
				_, err = v.CheckSec(context)
				if err != nil {
					return
				}
			}
			err = v.Exploit(context)
			return
		},
	}
}

func Vul2VulCmd(v vul.Vulnerability, alias []string, flagsCheckSec []cli.Flag, flagsExploit []cli.Flag, checkBeforeExploit bool) *cli.Command {
	checksec := Vul2ChecksecCmd(v, []string{"c"}, flagsCheckSec)
	checksec.Name = "checksec"
	checksec.Usage = "check vulnerability exists"

	exploit := Vul2ExploitCmd(v, []string{"x"}, flagsExploit, checkBeforeExploit)
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
