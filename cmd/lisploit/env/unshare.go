package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
	"sploit-spec/env"
)

var Unshare = &cli.Command{
	Name:    "Unshare",
	Aliases: []string{"m"},
	Usage:   "show whether support unshare command",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		result := env.Seccomp()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}
