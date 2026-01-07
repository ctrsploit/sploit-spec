package env

import (
	"context"
	"xsploit/env/auto"

	"github.com/urfave/cli/v3"
)

const (
	CommandNameAuto = "auto"
)

var (
	Auto = &cli.Command{
		Name:  CommandNameAuto,
		Usage: "auto",
		Action: func(context.Context, *cli.Command) (err error) {
			auto.Print()
			return
		},
	}
)
