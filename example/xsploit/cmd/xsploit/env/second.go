package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/example/xsploit/pkg/env"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
)

var Second = &cli.Command{
	Name:    "second",
	Aliases: []string{"s"},
	Usage:   "show the second info",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		result := env.Second()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}
