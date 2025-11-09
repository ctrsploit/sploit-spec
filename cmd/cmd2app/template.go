package main

// appTemplate is the Go source code template for the final executable.
// It creates a simple cli.App that wraps the target subcommand.
const appTemplate = `
package main

import (
	"log"
	"os"

	subcmd_pkg "{{.PkgPath}}"
	"github.com/urfave/cli/v3"
)

func main() {
	// The target variable (e.g., auto.Command) is already a pointer to cli.Command.
	targetCommand := subcmd_pkg.{{.VarName}}

	app := &cli.Command{
		Name:		targetCommand.Name,
		Usage:		targetCommand.Usage,
		Flags:		targetCommand.Flags,
		Action:		targetCommand.Action,
		Commands:	targetCommand.Commands, // Also include sub-subcommands if any
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
`
