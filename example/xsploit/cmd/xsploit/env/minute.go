package env

import (
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/urfave/cli/v2"
	"xsploit/env/minute"
)

var Minute = &cli.Command{
	Name:    "minute",
	Aliases: []string{"m"},
	Usage:   "show the minute info",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		minute.Print()
		return
	},
}
