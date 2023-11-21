package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/example/xsploit/pkg/env"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
)

const (
	CommandNameAuto = "auto"
)

type Result map[string]printer.Interface

var (
	Auto = &cli.Command{
		Name:  CommandNameAuto,
		Usage: "auto",
		Action: func(context *cli.Context) (err error) {
			s := env.Second()
			m := env.Minute()
			result := Result{
				"second": s,
				"minute": m,
			}
			fmt.Println(printer.Printer.Print(result))
			return
		},
	}
)
