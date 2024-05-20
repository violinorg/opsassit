package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
	"gopkg.in/yaml.v2"
)

func drainCmd() *cli.Command {
	return &cli.Command{
		Name:      "drain",
		Usage:     "Drain missing keys from the second YAML file into the first YAML file",
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

			updatedVars := drainYAML(vars1, vars2)
			updatedYAML, err := yaml.Marshal(updatedVars)
			if err != nil {
				fmt.Printf("Error marshalling updated YAML: %v\n", err)
				os.Exit(1)
			}

			err = ioutil.WriteFile(file1Path, updatedYAML, 0644)
			if err != nil {
				fmt.Printf("Error writing updated YAML to file1: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Successfully drained keys from file2 to file1.")
			return nil
		},
	}
}

func drainYAML(vars1, vars2 map[string]interface{}) map[string]interface{} {
	for key, val2 := range vars2 {
		if val1, exists := vars1[key]; exists {
			if val1 != val2 {
				vars1[key] = fmt.Sprintf("%v\" # from file2 = %v", val1, val2)
			}
		} else {
			vars1[key] = val2
		}
	}
	return vars1
}
