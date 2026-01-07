package env

import (
	"context"
	"xsploit/env/minute"

	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/urfave/cli/v3"
)

var Minute = &cli.Command{
	Name:    "minute",
	Aliases: []string{"m"},
	Usage:   "show the minute info",
	Action: func(context.Context, *cli.Command) (err error) {
		log.Logger.Debug("")
		minute.Print()
		return
	},
}
