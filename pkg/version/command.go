package version

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "version",
	Aliases: []string{},
	Usage:   "Show the sploit version information",
	Action: func(context.Context, *cli.Command) error {
		fmt.Println(DefaultVer())
		return nil
	},
}
