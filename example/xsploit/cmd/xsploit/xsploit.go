package main

import (
	"os"
	"xsploit/cmd/xsploit/auto"
	"xsploit/cmd/xsploit/checksec"
	"xsploit/cmd/xsploit/env"
	"xsploit/cmd/xsploit/exploit"
	"xsploit/cmd/xsploit/vul"

	"github.com/ctrsploit/sploit-spec/pkg/app"
	spec_version "github.com/ctrsploit/sploit-spec/pkg/spec-version"
	"github.com/ctrsploit/sploit-spec/pkg/version"
	"github.com/urfave/cli/v2"
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
			vul.Command,
			version.Command,
			spec_version.Command,
		},
	}
	app.InstallGlobalFlags(sploit)
	_ = sploit.Run(os.Args)
}
