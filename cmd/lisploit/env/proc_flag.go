package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
	"sploit-spec/env"
)

var ProcFlag = &cli.Command{
	Name:    "ProcFlag",
	Aliases: []string{"m"},
	Usage:   "show the Proc Protect Flag",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		result := env.Seccomp()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}
