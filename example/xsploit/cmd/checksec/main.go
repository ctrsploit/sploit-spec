package main

import (
	"os"
	"xsploit/cmd/xsploit/checksec"

	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

const (
	name = `x-sploit/checksec`
)

func main() {
	sploit := app.Command2App(checksec.Command)
	sploit.Name = name
	app.InstallGlobalFlags(sploit)
	err := sploit.Run(os.Args)
	if err != nil {
		awesome_error.CheckFatal(err)
	}
}
