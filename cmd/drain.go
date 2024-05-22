package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
	"os"
	"strings"
)

func drainCmd() *cli.Command {
	return &cli.Command{
		Name:      "drain",
		Usage:     "Drain missing keys from the second YAML file into the first YAML file",
		ArgsUsage: "[file1] [file2]",
		Action: func(c *cli.Context) error {
			file1Path := os.Getenv("FILE1_PATH")
			file2Path := os.Getenv("FILE2_PATH")

			if file1Path == "" || file2Path == "" {
				if c.NArg() != 2 {
					return fmt.Errorf("expected exactly 2 arguments")
				}
				file1Path = c.Args().Get(0)
				file2Path = c.Args().Get(1)
			}

			approved := c.Bool("approved")

			// Read file1 to check for "# OpsAssist Verified"
			file1Content, err := os.ReadFile(file1Path)
			if err != nil {
				return fmt.Errorf("error reading file1: %v", err)
			}

			if strings.Contains(string(file1Content), "# OpsAssist Verified") {
				fmt.Println("The file is already tuned.")
				return nil
			}

			vars1, err := actions.LoadVariablesFromYAMLWithOrder(file1Path)
			if err != nil {
				return fmt.Errorf("error loading file1: %v", err)
			}

			vars2, err := actions.LoadVariablesFromYAMLWithOrder(file2Path)
			if err != nil {
				return fmt.Errorf("error loading file2: %v", err)
			}

			updatedYAML, err := actions.GenerateUpdatedYAML(vars1, vars2)
			if err != nil {
				return fmt.Errorf("error generating updated YAML: %v", err)
			}

			// Output the changes
			fmt.Println("Announcing changes:")
			for _, key := range vars1.Keys {
				val1 := vars1.Values[key]
				if val2, exists := vars2.Values[key]; exists && !actions.ValuesEqual(val1, val2) {
					color.New(color.FgGreen).Printf("# Added from file2 - %s: %v\n", key, val2)
					fmt.Printf("%s: %v\n", key, val1)
				} else {
					fmt.Printf("%s: %v\n", key, val1)
				}
			}
			for _, key := range vars2.Keys {
				if _, exists := vars1.Values[key]; !exists {
					color.New(color.FgGreen).Printf("%s: %v\n", key, vars2.Values[key])
				}
			}

			if approved {
				updatedYAML += "\n# OpsAssist Verified\n"

				err = os.WriteFile(file1Path, []byte(updatedYAML), 0644)
				if err != nil {
					return fmt.Errorf("error writing updated YAML to file1: %v", err)
				}

				fmt.Println("Successfully drained keys from file2 to file1.")
			} else {
				fmt.Println("Run the command with --approved to apply these changes.")
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "approved",
				Usage: "Apply the changes to the file",
			},
		},
	}
}
