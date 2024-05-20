package cmd

import (
	"github.com/urfave/cli/v2"
	"os"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func addGitLabFlags(flags []cli.Flag) []cli.Flag {
	return append(flags,
		&cli.StringFlag{
			Name:    "gitlab-url",
			Usage:   "GitLab URL",
			Value:   getEnvOrDefault("OA_GITLAB_URL", ""),
			EnvVars: []string{"OA_GITLAB_URL"},
		},
		&cli.StringFlag{
			Name:    "gitlab-token",
			Usage:   "GitLab personal access token",
			Value:   getEnvOrDefault("OA_GITLAB_TOKEN", ""),
			EnvVars: []string{"OA_GITLAB_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "project-id",
			Usage:   "GitLab project ID",
			Value:   getEnvOrDefault("OA_GITLAB_PROJECT_ID", ""),
			EnvVars: []string{"OA_GITLAB_PROJECT_ID"},
		},
		&cli.StringFlag{
			Name:    "base-branch",
			Usage:   "Base branch for the new branch",
			Value:   getEnvOrDefault("OA_GITLAB_BASE_BRANCH", ""),
			EnvVars: []string{"OA_GITLAB_BASE_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "new-branch",
			Usage:   "Name of the new branch",
			Value:   getEnvOrDefault("OA_GITLAB_NEW_BRANCH", ""),
			EnvVars: []string{"OA_GITLAB_NEW_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "target-branch",
			Usage:   "Target branch for the merge request",
			Value:   getEnvOrDefault("OA_GITLAB_TARGET_BRANCH", "$OA_GITLAB_BASE_BRANCH"),
			EnvVars: []string{"OA_GITLAB_TARGET_BRANCH"},
		},
		&cli.BoolFlag{
			Name:    "auto-mr",
			Usage:   "Automatically create a merge request in GitLab",
			EnvVars: []string{"OA_GITLAB_AUTO_MR"},
		},
	)
}
