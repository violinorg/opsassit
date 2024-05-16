package cmd

import (
	"fmt"
)

var commands = []*cli.Command{
	DiffValuesCmd(),
	DiffKeysCmd(),
}

func CreateApp() *cli.App {
	c := cli.NewApp()
	c.EnableBashCompletion = true
	c.Usage = "Compare YAML files"
	c.Commands = commands
	c.CommandNotFound = command404

	return c
}

// CommandNotFoundError is returned when CLI command is not found.
type CommandNotFoundError struct {
	Command string
}

func (e CommandNotFoundError) Error() string {
	return fmt.Sprintf("👻 Command %q not found", e.Command)
}

func command404(_ *cli.Context, s string) {
	err := CommandNotFoundError{
		Command: s,
	}
	panic(err)
}
