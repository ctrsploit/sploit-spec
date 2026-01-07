package env

import "github.com/urfave/cli/v3"

var Command = &cli.Command{
	Name:    "env",
	Aliases: []string{"e"},
	Usage:   "Collect information",
	Commands: []*cli.Command{
		Auto,
		Second,
		Minute,
		Upload,
	},
}
