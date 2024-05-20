package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
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

			vars1, order1, err := actions.LoadVariablesFromYAMLWithOrder(file1Path)
			if err != nil {
				fmt.Printf("Error loading file1: %v\n", err)
				os.Exit(1)
			}

			vars2, _, err := actions.LoadVariablesFromYAMLWithOrder(file2Path)
			if err != nil {
				fmt.Printf("Error loading file2: %v\n", err)
				os.Exit(1)
			}

			updatedYAML, err := generateUpdatedYAML(vars1, vars2, order1)
			if err != nil {
				fmt.Printf("Error generating updated YAML: %v\n", err)
				os.Exit(1)
			}

			err = ioutil.WriteFile(file1Path, []byte(updatedYAML), 0644)
			if err != nil {
				fmt.Printf("Error writing updated YAML to file1: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Successfully drained keys from file2 to file1.")
			return nil
		},
	}
}

func generateUpdatedYAML(vars1, vars2 map[string]interface{}, order1 []string) (string, error) {
	var builder strings.Builder
	builder.WriteString("---\n")

	// Preserve the order of keys from file1
	for _, key := range order1 {
		val1 := vars1[key]
		if val2, exists := vars2[key]; exists && !reflect.DeepEqual(val1, val2) {
			builder.WriteString(fmt.Sprintf("# from file2 - %s: %v\n", key, val2))
		}
		builder.WriteString(fmt.Sprintf("%s: %v\n", key, val1))
	}

	// Add keys from file2 that do not exist in file1
	for key, val2 := range vars2 {
		if _, exists := vars1[key]; !exists {
			builder.WriteString(fmt.Sprintf("%s: %v\n", key, val2))
		}
	}

	return builder.String(), nil
}
