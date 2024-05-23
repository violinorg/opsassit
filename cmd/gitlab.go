package cmd

import (
	"fmt"
	"os"

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
		ArgsUsage: "[filePath] [baseBranch] [newBranch] [targetBranch] [projectID]",
		Action:    autoMrAction,
		Flags:     addGitLabFlags(nil),
	}
}

func autoMrAction(c *cli.Context) error {
	filePath := os.Getenv("OA_GITLAB_AUTOMR_FILE_PATH")
	gitlabURL := os.Getenv("OA_GITLAB_URL")
	gitlabToken := os.Getenv("OA_GITLAB_TOKEN")
	projectID := os.Getenv("OA_GITLAB_PROJECT_ID")
	baseBranch := os.Getenv("OA_GITLAB_BASE_BRANCH")
	newBranch := os.Getenv("OA_GITLAB_NEW_BRANCH")
	targetBranch := os.Getenv("OA_GITLAB_TARGET_BRANCH")

	if filePath == "" {
		filePath = c.Args().Get(0)
	}
	if gitlabURL == "" {
		gitlabURL = c.String("gitlab-url")
	}
	if gitlabToken == "" {
		gitlabToken = c.String("gitlab-token")
	}
	if projectID == "" {
		projectID = c.String("project-id")
	}
	if baseBranch == "" {
		baseBranch = c.String("base-branch")
	}
	if newBranch == "" {
		newBranch = c.String("new-branch")
	}
	if targetBranch == "" {
		targetBranch = c.String("target-branch")
	}

	err := actions.HandleGitLabMergeRequest(gitlabURL, gitlabToken, filePath, baseBranch, newBranch, targetBranch, projectID)
	if err != nil {
		return fmt.Errorf("error handling GitLab merge request: %v", err)
	}

	return nil
}
