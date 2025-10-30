package spec_version

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const version = "v0.7.0"

var Command = &cli.Command{
	Name:    "spec-version",
	Aliases: []string{},
	Usage:   "show which version of spec the sploit tool follows",
	Action: func(context *cli.Context) error {
		fmt.Println(version)
		return nil
	},
}
