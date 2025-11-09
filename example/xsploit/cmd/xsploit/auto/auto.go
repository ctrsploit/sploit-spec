package auto

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

const (
	CommandNameAuto = "auto"
)

var (
	Command = &cli.Command{
		Name:    CommandNameAuto,
		Usage:   "auto gathering information, detect vulnerabilities and run exploits",
		Aliases: []string{"a"},
		Action: func(context.Context, *cli.Command) (err error) {
			fmt.Println("TODO")
			return
		},
	}
)
