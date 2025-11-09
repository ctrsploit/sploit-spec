package app

import (
	"context"

	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/urfave/cli/v3"
)

func Vul2ChecksecCmd(v vul.Vulnerability, alias []string, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:    v.GetName(),
		Aliases: alias,
		Usage:   v.GetDescription(),
		Flags:   flags,
		Action: func(ctx context.Context, cmd *cli.Command) (err error) {
			for _, flag := range cmd.Flags {
				name := flag.Names()[0]
				val := cmd.Value(name)
				ctx = context.WithValue(ctx, name, val)
			}
			_, err = v.CheckSec(ctx)
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
		Action: func(ctx context.Context, cmd *cli.Command) (err error) {
			for _, flag := range cmd.Flags {
				name := flag.Names()[0]
				val := cmd.Value(name)
				ctx = context.WithValue(ctx, name, val)
			}
			if checkBeforeExploit {
				_, err = v.CheckSec(ctx)
				if err != nil {
					return
				}
			}
			err = v.Exploit(ctx)
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
		Commands: []*cli.Command{
			checksec,
			exploit,
		},
	}
}
