package cmd

import "github.com/urfave/cli/v2"

func addGitLabFlags(flags []cli.Flag) []cli.Flag {
	return append(flags,
		&cli.StringFlag{
			Name:    "gitlab-url",
			Usage:   "GitLab instance URL",
			EnvVars: []string{"OA_GITLAB_URL"},
		},
		&cli.StringFlag{
			Name:    "gitlab-token",
			Usage:   "GitLab access token",
			EnvVars: []string{"OA_GITLAB_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "gitlab-project-id",
			Usage:   "GitLab project ID",
			EnvVars: []string{"OA_GITLAB_PROJECT_ID"},
		},
		&cli.StringFlag{
			Name:    "mr-base-branch",
			Usage:   "Base branch for the new branch",
			EnvVars: []string{"OA_GITLAB_MR_BASE_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "mr-new-branch",
			Usage:   "New branch name",
			EnvVars: []string{"OA_GITLAB_MR_NEW_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "mr-target-branch",
			Usage:   "Target branch for the merge request",
			EnvVars: []string{"OA_GITLAB_MR_TARGET_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "mr-title",
			Usage:   "Title for the merge request",
			Value:   "Draft: OpsAssist auto-mr",
			EnvVars: []string{"OA_GITLAB_MR_TITLE"},
		},
		&cli.StringFlag{
			Name:    "mr-description",
			Usage:   "Description for the merge request",
			Value:   "OpsAssist description",
			EnvVars: []string{"OA_GITLAB_MR_DESCRIPTION"},
		},
		&cli.BoolFlag{
			Name:    "mr-squash",
			Usage:   "Squash commits in the merge request",
			Value:   true,
			EnvVars: []string{"OA_GITLAB_MR_SQUASH"},
		},
		&cli.BoolFlag{
			Name:    "mr-remove-source-branch",
			Usage:   "Remove source branch after merge",
			Value:   true,
			EnvVars: []string{"OA_GITLAB_MR_REMOVE_SOURCE_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "src-file-path",
			Usage:   "Source file path to read content from",
			EnvVars: []string{"OA_GITLAB_SRC_FILE_PATH"},
		},
		&cli.StringFlag{
			Name:    "file-path",
			Usage:   "Destination file path in the repository",
			EnvVars: []string{"OA_GITLAB_FILE_PATH"},
		},
	)
}
