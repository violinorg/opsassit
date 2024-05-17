package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
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
			autoMR := c.Bool("auto-mr")

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
			resultFilePath := "comparison_result.md"

			err = actions.SaveComparisonResult(resultFilePath, onlyInFile1, onlyInFile2)
			if err != nil {
				fmt.Printf("Error saving comparison result: %v\n", err)
				os.Exit(1)
			}

			if autoMR {
				gitlabURL := c.String("gitlab-url")
				gitlabToken := c.String("gitlab-token")
				projectID := c.Int("project-id")
				baseBranch := c.String("base-branch")
				newBranch := c.String("new-branch")
				targetBranch := c.String("target-branch")

				gitlabClient, err := actions.NewGitLabClient(gitlabURL, gitlabToken)
				if err != nil {
					fmt.Printf("Error creating GitLab client: %v\n", err)
					os.Exit(1)
				}

				err = gitlabClient.CreateBranch(projectID, newBranch, baseBranch)
				if err != nil {
					fmt.Printf("Error creating branch: %v\n", err)
					os.Exit(1)
				}

				content, err := os.ReadFile(resultFilePath)
				if err != nil {
					fmt.Printf("Error reading result file: %v\n", err)
					os.Exit(1)
				}

				err = gitlabClient.CreateFile(projectID, newBranch, resultFilePath, string(content))
				if err != nil {
					fmt.Printf("Error creating file: %v\n", err)
					os.Exit(1)
				}

				err = gitlabClient.CreateMergeRequest(projectID, newBranch, targetBranch, "WIP: Comparison Result")
				if err != nil {
					fmt.Printf("Error creating merge request: %v\n", err)
					os.Exit(1)
				}
			} else {
				fmt.Println("Keys only in file1:")
				for _, key := range onlyInFile1 {
					fmt.Println(key)
				}

				fmt.Println("\nKeys only in file2:")
				for _, key := range onlyInFile2 {
					fmt.Println(key)
				}
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "gitlab-url",
				Usage: "GitLab URL",
			},
			&cli.StringFlag{
				Name:  "gitlab-token",
				Usage: "GitLab personal access token",
			},
			&cli.IntFlag{
				Name:  "project-id",
				Usage: "GitLab project ID",
			},
			&cli.StringFlag{
				Name:  "base-branch",
				Usage: "Base branch for the new branch",
			},
			&cli.StringFlag{
				Name:  "new-branch",
				Usage: "Name of the new branch",
			},
			&cli.StringFlag{
				Name:  "target-branch",
				Usage: "Target branch for the merge request",
			},
			&cli.BoolFlag{
				Name:  "auto-mr",
				Usage: "Automatically create a merge request in GitLab",
			},
		},
	}
}
