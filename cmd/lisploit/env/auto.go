package env

import (
	"github.com/urfave/cli/v2"
	"xsploit/env"
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
