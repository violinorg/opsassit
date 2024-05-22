package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "OpsAssist"
	app.Usage = "A CLI tool for YAML file operations"
	app.Commands = []*cli.Command{
		drainCmd(),
		diffValuesCmd(),
	}
	return app
}

// CommandNotFoundError is returned when CLI command is not found.
type CommandNotFoundError struct {
	Command string
}

func (e CommandNotFoundError) Error() string {
	return fmt.Sprintf("👻 Command %q not found", e.Command)
}

func command404(c *cli.Context, s string) {
	err := CommandNotFoundError{
		Command: s,
	}
	panic(err)
}

//package cmd
//
//import (
//	"fmt"
//
//	"github.com/urfave/cli/v2"
//)
//
//var commands = []*cli.Command{
//	diffValuesCmd(),
//	diffKeysCmd(),
//	drainCmd(),
//}
//
//func CreateApp() *cli.App {
//	c := cli.NewApp()
//	c.EnableBashCompletion = true
//	c.Usage = "Compare YAML files"
//	c.Commands = commands
//	c.CommandNotFound = command404
//
//	return c
//}
//
