package env

import "github.com/urfave/cli/v2"

var Command = &cli.Command{
	Name:    "env",
	Aliases: []string{"e"},
	Usage:   "Collect information",
	Subcommands: []*cli.Command{
		Auto,
		Second,
		Minute,
	},
}
