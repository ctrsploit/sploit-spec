package main

import (
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/urfave/cli/v2"
	"os"
)

const usage = `An example sploit tool follows sploit-spec`

func main() {
	sploit := &cli.App{
		Name:     "xsploit",
		Usage:    usage,
		Commands: []*cli.Command{},
	}
	app.InstallGlobalFlags(sploit)
	_ = sploit.Run(os.Args)
}
