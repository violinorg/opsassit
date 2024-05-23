package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
)

func gitlabCmd() *cli.Command {
	return &cli.Command{
		Name:  "gitlab",
		Usage: "Commands for working with GitLab",
		Subcommands: []*cli.Command{
			autoMrCmd(),
		},
	}
}

func autoMrCmd() *cli.Command {
	return &cli.Command{
		Name:      "auto-mr",
		Usage:     "Automatically create a merge request",
		ArgsUsage: "[filePath]",
		Action:    autoMrAction,
		Flags:     addGitLabFlags(nil),
	}
}

func autoMrAction(c *cli.Context) error {
	filePath := c.Args().First()

	gitlabURL := c.String("gitlab-url")
	gitlabToken := c.String("gitlab-token")
	projectID := c.String("project-id")
	baseBranch := c.String("base-branch")
	newBranch := c.String("new-branch")
	targetBranch := c.String("target-branch")

	missingParameters := []string{}
	if filePath == "" {
		missingParameters = append(missingParameters, "filePath")
	}
	if gitlabURL == "" {
		missingParameters = append(missingParameters, "gitlab-url")
	}
	if gitlabToken == "" {
		missingParameters = append(missingParameters, "gitlab-token")
	}
	if projectID == "" {
		missingParameters = append(missingParameters, "project-id")
	}
	if baseBranch == "" {
		missingParameters = append(missingParameters, "base-branch")
	}
	if newBranch == "" {
		missingParameters = append(missingParameters, "new-branch")
	}
	if targetBranch == "" {
		missingParameters = append(missingParameters, "target-branch")
	}

	if len(missingParameters) > 0 {
		return fmt.Errorf("missing required parameters: %v", missingParameters)
	}

	err := actions.HandleGitLabMergeRequest(gitlabURL, gitlabToken, filePath, baseBranch, newBranch, targetBranch, projectID)
	if err != nil {
		return fmt.Errorf("error handling GitLab merge request: %v", err)
	}

	return nil
}
