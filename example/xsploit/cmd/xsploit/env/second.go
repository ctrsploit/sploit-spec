package env

import (
	"context"
	"xsploit/env/second"

	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/urfave/cli/v3"
)

var Second = &cli.Command{
	Name:    "second",
	Aliases: []string{"s"},
	Usage:   "show the second info",
	Action: func(context.Context, *cli.Command) (err error) {
		log.Logger.Debug("")
		second.Print()
		return
	},
}
