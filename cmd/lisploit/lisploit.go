package main

import (
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ctrsploit/sploit-spec/pkg/version"
	"github.com/urfave/cli/v2"
	"os"
	"cmd/lisploit/auto"
	"cmd/lisploit/checksec"
	"cmd/lisploit/env"
	"cmd/lisploit/exploit"
	"cmd/lisploit/vul"
)

const usage = `linux kernel sploit tool follows sploit-spec`

func init() {
	version.ProductName = "lisploit"
}

func main() {
	sploit := &cli.App{
		Name:  "lisploit",
		Usage: usage,
		Commands: []*cli.Command{
			auto.Command,
			env.Command,
			checksec.Command,
			exploit.Command,
			vul.Command,
			version.Command,
		},
	}
	app.InstallGlobalFlags(sploit)
	_ = sploit.Run(os.Args)
}
