package cmd

import (
	"ConfigGuard/actions"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func DiffValuesCmd() *cli.Command {
	return &cli.Command{
		Name:      "values",
		Usage:     "Compare the values of two YAML files",
		ArgsUsage: "[file1] [file2]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return fmt.Errorf("expected exactly 2 arguments")
			}

			file1Path := c.Args().Get(0)
			file2Path := c.Args().Get(1)

			vars1, err := actions.LoadVariablesFromYAML(file1Path)
			if err != nil {
				fmt.Printf("Error loading file1: %v\n", err)
				os.Exit(1)
			}

			vars2, err := actions.LoadVariablesFromYAML(file2Path)
			if err != nil {
				fmt.Printf("Error loading file2: %v\n", err)
				os.Exit(1)
			}

			differences := actions.CompareValues(vars1, vars2)

			if len(differences) > 0 {
				fmt.Println("Differences found:")
				for key, vals := range differences {
					fmt.Printf("Variable %s: %v (file1) != %v (file2)\n", key, vals[0], vals[1])
				}
			} else {
				fmt.Println("No differences found.")
			}

			return nil
		},
	}
}
