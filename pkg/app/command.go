package app

import (
	"fmt"

	"github.com/ctrsploit/sploit-spec/pkg/version"
	"github.com/urfave/cli/v3"
)

func Command2App(command *cli.Command) (app *cli.Command) {
	return &cli.Command{
		Name:     command.Name,
		Usage:    command.Usage,
		Action:   command.Action,
		Commands: append(command.Commands, version.Command),
		Flags:    command.Flags,
		Before:   command.Before,
	}
}

// ReplaceSubCommand searches for a subcommand by its name or alias within a parent command's
// subcommands list and replaces it with a new subcommand.
//
// Parameters:
//   - parentCmd: The parent command whose subcommands list will be modified. Cannot be nil.
//   - nameOrAlias: The name or an alias of the subcommand to be replaced.
//   - newSubCmd: The new subcommand to insert. Cannot be nil.
//
// Returns an error if the parent command or new subcommand is nil, or if no subcommand
// with the given name or alias is found. Otherwise, it returns nil upon successful replacement.
func ReplaceSubCommand(commandName string, command *cli.Command, newSubCommand *cli.Command) error {
	if command == nil {
		return fmt.Errorf("parent command cannot be nil")
	}
	if newSubCommand == nil {
		return fmt.Errorf("new subcommand cannot be nil")
	}
	found := false
	for i, subcommand := range command.Commands {
		if subcommand.Name == commandName {
			command.Commands[i] = newSubCommand
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("subcommand '%s' not found in command '%s'", commandName, command.Name)
	}
	return nil
}
