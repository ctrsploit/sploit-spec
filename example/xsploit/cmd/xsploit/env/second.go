package env

import (
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/urfave/cli/v2"
	"xsploit/env/second"
)

var Second = &cli.Command{
	Name:    "second",
	Aliases: []string{"s"},
	Usage:   "show the second info",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		second.Print()
		return
	},
}
