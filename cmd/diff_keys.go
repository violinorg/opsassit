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
			gitlabToken := c.String("gitlab-token")
			projectID := c.Int("project-id")
			baseBranch := c.String("base-branch")
			newBranch := c.String("new-branch")
			targetBranch := c.String("target-branch")

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

			gitlabClient, err := actions.NewGitLabClient(gitlabToken)
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

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "gitlab-token",
				Usage:    "GitLab personal access token",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "project-id",
				Usage:    "GitLab project ID",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "base-branch",
				Usage:    "Base branch for the new branch",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "new-branch",
				Usage:    "Name of the new branch",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "target-branch",
				Usage:    "Target branch for the merge request",
				Required: true,
			},
		},
	}
}
