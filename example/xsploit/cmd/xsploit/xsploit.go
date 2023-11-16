package main

import (
	"github.com/ctrsploit/sploit-spec/example/xsploit/cmd/xsploit/auto"
	"github.com/ctrsploit/sploit-spec/example/xsploit/cmd/xsploit/checksec"
	"github.com/ctrsploit/sploit-spec/example/xsploit/cmd/xsploit/env"
	"github.com/ctrsploit/sploit-spec/example/xsploit/cmd/xsploit/exploit"
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ctrsploit/sploit-spec/pkg/version"
	"github.com/urfave/cli/v2"
	"os"
)

const usage = `An example sploit tool follows sploit-spec`

func init() {
	version.ProductName = "xsploit"
}

func main() {
	sploit := &cli.App{
		Name:  "xsploit",
		Usage: usage,
		Commands: []*cli.Command{
			auto.Command,
			env.Command,
			checksec.Command,
			exploit.Command,
			version.Command,
		},
	}
	app.InstallGlobalFlags(sploit)
	_ = sploit.Run(os.Args)
}
