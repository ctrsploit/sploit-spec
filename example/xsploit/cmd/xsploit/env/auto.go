package env

import (
	"github.com/ctrsploit/sploit-spec/example/xsploit/env"
	"github.com/urfave/cli/v2"
)

const (
	CommandNameAuto = "auto"
)

var (
	Auto = &cli.Command{
		Name:  CommandNameAuto,
		Usage: "auto",
		Action: func(context *cli.Context) (err error) {
			env.Auto()
			return
		},
	}
)
