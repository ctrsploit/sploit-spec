package spec_version

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

const version = "v0.7.0"

var Command = &cli.Command{
	Name:    "spec-version",
	Aliases: []string{},
	Usage:   "show which version of spec the sploit tool follows",
	Action: func(context.Context, *cli.Command) error {
		fmt.Println(version)
		return nil
	},
}
