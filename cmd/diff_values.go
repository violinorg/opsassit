package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
)

func diffValuesCmd() *cli.Command {
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
			autoMR := c.Bool("auto-mr")

			vars1, err := actions.LoadVariablesFromYAMLWithOrder(file1Path)
			if err != nil {
				return fmt.Errorf("error loading file1: %v", err)
			}

			vars2, err := actions.LoadVariablesFromYAMLWithOrder(file2Path)
			if err != nil {
				return fmt.Errorf("error loading file2: %v", err)
			}

			differences := actions.CompareValues(vars1, vars2)
			resultFilePath := "comparison_values_result.md"

			err = actions.SaveValuesComparisonResult(resultFilePath, differences)
			if err != nil {
				return fmt.Errorf("error saving comparison result: %v", err)
			}

			if autoMR {
				gitlabURL := c.String("gitlab-url")
				gitlabToken := c.String("gitlab-token")
				projectID := c.String("project-id")
				baseBranch := c.String("base-branch")
				newBranch := c.String("new-branch")
				targetBranch := c.String("target-branch")

				err = actions.HandleGitLabMergeRequest(gitlabURL, gitlabToken, resultFilePath, baseBranch, newBranch, targetBranch, projectID)
				if err != nil {
					return fmt.Errorf("error handling GitLab merge request: %v", err)
				}
			} else {
				fmt.Println("Differences in values:")
				for key, vals := range differences {
					fmt.Printf("%s: %v -> %v\n", key, vals[0], vals[1])
				}
			}

			fmt.Println("Comparison completed successfully.")

			return nil
		},
		Flags: addGitLabFlags(nil),
	}
}
