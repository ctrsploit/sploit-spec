package auto

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const (
	CommandNameAuto = "auto"
)

var (
	Command = &cli.Command{
		Name:  CommandNameAuto,
		Usage: "auto",
		Action: func(context *cli.Context) (err error) {
			fmt.Println("TODO")
			return
		},
	}
)
