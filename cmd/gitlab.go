package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
)

func GitlabCmd() *cli.Command {
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
	// Сначала получаем значения из переменных окружения
	filePath := os.Getenv("OA_GITLAB_AUTOMR_FILE_PATH")
	gitlabURL := os.Getenv("OA_GITLAB_URL")
	gitlabToken := os.Getenv("OA_GITLAB_TOKEN")
	projectID := os.Getenv("OA_GITLAB_PROJECT_ID")
	baseBranch := os.Getenv("OA_GITLAB_BASE_BRANCH")
	newBranch := os.Getenv("OA_GITLAB_NEW_BRANCH")
	targetBranch := os.Getenv("OA_GITLAB_TARGET_BRANCH")

	// Затем переопределяем их значениями из флагов, если они указаны
	if c.Args().Get(0) != "" {
		filePath = c.Args().Get(0)
	}
	if c.String("gitlab-url") != "" {
		gitlabURL = c.String("gitlab-url")
	}
	if c.String("gitlab-token") != "" {
		gitlabToken = c.String("gitlab-token")
	}
	if c.String("project-id") != "" {
		projectID = c.String("project-id")
	}
	if c.String("base-branch") != "" {
		baseBranch = c.String("base-branch")
	}
	if c.String("new-branch") != "" {
		newBranch = c.String("new-branch")
	}
	if c.String("target-branch") != "" {
		targetBranch = c.String("target-branch")
	}

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
