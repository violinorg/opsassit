package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
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

			autoMR := c.Bool("auto-mr")

			// Read file1 to check for "# OpsAssist Verified"
			file1Content, err := os.ReadFile(file1Path)
			if err != nil {
				return fmt.Errorf("error reading file1: %v", err)
			}

			if strings.Contains(string(file1Content), "# OpsAssist Verified") {
				color.Green("The file is already tuned.")
				return nil
			}

			vars1, order1, err := actions.LoadVariablesFromYAMLWithOrder(file1Path)
			if err != nil {
				return fmt.Errorf("error loading file1: %v", err)
			}

			vars2, _, err := actions.LoadVariablesFromYAMLWithOrder(file2Path)
			if err != nil {
				return fmt.Errorf("error loading file2: %v", err)
			}

			updatedYAML, err := generateUpdatedYAML(vars1, vars2, order1)
			if err != nil {
				return fmt.Errorf("error generating updated YAML: %v", err)
			}

			updatedYAML += "\n# OpsAssist Verified\n"

			err = os.WriteFile(file1Path, []byte(updatedYAML), 0644)
			if err != nil {
				return fmt.Errorf("error writing updated YAML to file1: %v", err)
			}

			color.Green("Successfully drained keys from file2 to file1.")

			if autoMR {
				gitlabURL := c.String("gitlab-url")
				gitlabToken := c.String("gitlab-token")
				projectID := c.String("project-id")
				baseBranch := c.String("base-branch")
				newBranch := c.String("new-branch")
				targetBranch := c.String("target-branch")

				err = actions.HandleGitLabMergeRequest(gitlabURL, gitlabToken, file1Path, baseBranch, newBranch, targetBranch, projectID)
				if err != nil {
					return fmt.Errorf("error handling GitLab merge request: %v", err)
				}
			}

			return nil
		},
		Flags: addGitLabFlags(nil),
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
