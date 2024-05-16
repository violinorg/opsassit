package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
	"os"
)

func diffKeysCmd() *cli.Command {
	return &cli.Command{
		Name:      "keys",
		Usage:     "Compare the keys of two YAML files",
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

			onlyInFile1, onlyInFile2 := actions.CompareKeys(vars1, vars2)

			if len(onlyInFile1) > 0 {
				fmt.Println("Keys only in file1:")
				for _, key := range onlyInFile1 {
					fmt.Println(key)
				}
			} else {
				fmt.Println("No keys unique to file1.")
			}

			if len(onlyInFile2) > 0 {
				fmt.Println("Keys only in file2:")
				for _, key := range onlyInFile2 {
					fmt.Println(key)
				}
			} else {
				fmt.Println("No keys unique to file2.")
			}

			return nil
		},
	}
}
