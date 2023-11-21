package app

import (
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/urfave/cli/v2"
)

func Vul2Cmd(v vul.Vulnerability) *cli.Command {
	return &cli.Command{
		Name:  v.GetName(),
		Usage: v.GetDescription(),
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
