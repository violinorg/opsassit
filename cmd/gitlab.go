package cmd

import (
	"fmt"
	"os"
	"strconv"

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
	// Установим значения флагов и переменных окружения
	gitlabURL := c.String("gitlab-url")
	if gitlabURL == "" {
		gitlabURL = os.Getenv("OA_GITLAB_URL")
	}
	if gitlabURL == "" {
		return fmt.Errorf("gitlab URL must be provided as argument or environment variable")
	}

	gitlabToken := c.String("gitlab-token")
	if gitlabToken == "" {
		gitlabToken = os.Getenv("OA_GITLAB_TOKEN")
	}
	if gitlabToken == "" {
		return fmt.Errorf("gitlab token must be provided as argument or environment variable")
	}

	projectID := c.String("gitlab-project-id")
	if projectID == "" {
		projectID = os.Getenv("OA_GITLAB_PROJECT_ID")
	}
	if projectID == "" {
		return fmt.Errorf("project ID must be provided as argument or environment variable")
	}

	baseBranch := c.String("mr-base-branch")
	if baseBranch == "" {
		baseBranch = os.Getenv("OA_GITLAB_MR_BASE_BRANCH")
	}
	if baseBranch == "" {
		return fmt.Errorf("base branch must be provided as argument or environment variable")
	}

	newBranch := c.String("mr-new-branch")
	if newBranch == "" {
		newBranch = os.Getenv("OA_GITLAB_MR_NEW_BRANCH")
	}
	if newBranch == "" {
		return fmt.Errorf("new branch must be provided as argument or environment variable")
	}

	targetBranch := c.String("mr-target-branch")
	if targetBranch == "" {
		targetBranch = os.Getenv("OA_GITLAB_MR_TARGET_BRANCH")
	}
	if targetBranch == "" {
		return fmt.Errorf("target branch must be provided as argument or environment variable")
	}

	mrTitle := c.String("mr-title")
	if mrTitle == "" {
		mrTitle = os.Getenv("OA_GITLAB_MR_TITLE")
	}
	if mrTitle == "" {
		return fmt.Errorf("merge request title must be provided as argument or environment variable")
	}

	mrDescription := c.String("mr-description")
	if mrDescription == "" {
		mrDescription = os.Getenv("OA_GITLAB_MR_DESCRIPTION")
	}
	if mrDescription == "" {
		return fmt.Errorf("merge request description must be provided as argument or environment variable")
	}

	mrSquashStr := c.String("mr-squash")
	if mrSquashStr == "" {
		mrSquashStr = os.Getenv("OA_GITLAB_MR_SQUASH")
	}
	mrSquash, err := strconv.ParseBool(mrSquashStr)
	if err != nil {
		mrSquash = c.Bool("mr-squash")
	}

	mrRemoveSourceBranchStr := c.String("mr-remove-source-branch")
	if mrRemoveSourceBranchStr == "" {
		mrRemoveSourceBranchStr = os.Getenv("OA_GITLAB_MR_REMOVE_SOURCE_BRANCH")
	}
	mrRemoveSourceBranch, err := strconv.ParseBool(mrRemoveSourceBranchStr)
	if err != nil {
		mrRemoveSourceBranch = c.Bool("mr-remove-source-branch")
	}

	srcFilePath := c.String("src-file-path")
	if srcFilePath == "" {
		srcFilePath = os.Getenv("OA_GITLAB_SRC_FILE_PATH")
	}
	if srcFilePath == "" {
		return fmt.Errorf("source file path must be provided as argument or environment variable")
	}

	filePath := c.String("file-path")
	if filePath == "" {
		filePath = os.Getenv("OA_GITLAB_FILE_PATH")
	}
	if filePath == "" {
		return fmt.Errorf("file path must be provided as argument or environment variable")
	}

	err = actions.HandleGitLabMergeRequest(
		gitlabURL,
		gitlabToken,
		srcFilePath,
		filePath,
		baseBranch,
		newBranch,
		targetBranch,
		projectID,
		mrTitle,
		mrDescription,
		mrSquash,
		mrRemoveSourceBranch,
	)
	if err != nil {
		return fmt.Errorf("error handling GitLab merge request: %v", err)
	}

	return nil
}
