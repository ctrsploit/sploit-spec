package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
	"xsploit/env"
)

var Minute = &cli.Command{
	Name:    "minute",
	Aliases: []string{"m"},
	Usage:   "show the minute info",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		result := env.Minute()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}
