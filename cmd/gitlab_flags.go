package cmd

import (
	"github.com/urfave/cli/v2"
)

//func getEnvOrDefault(key, defaultValue string) string {
//	if value, exists := os.LookupEnv(key); exists {
//		return value
//	}
//	return defaultValue
//}

func addGitLabFlags(flags []cli.Flag) []cli.Flag {
	return append(flags,
		&cli.StringFlag{
			Name:  "gitlab-url",
			Usage: "GitLab URL",
			//EnvVars: []string{"OA_GITLAB_URL"},
		},
		&cli.StringFlag{
			Name:  "gitlab-token",
			Usage: "GitLab token",
			//EnvVars: []string{"OA_GITLAB_TOKEN"},
		},
		&cli.StringFlag{
			Name:  "project-id",
			Usage: "GitLab project ID",
			//EnvVars: []string{"OA_GITLAB_PROJECT_ID"},
		},
		&cli.StringFlag{
			Name:  "base-branch",
			Usage: "GitLab base branch",
			//EnvVars: []string{"OA_GITLAB_BASE_BRANCH"},
		},
		&cli.StringFlag{
			Name:  "new-branch",
			Usage: "GitLab new branch",
			//EnvVars: []string{"OA_GITLAB_NEW_BRANCH"},
		},
		&cli.StringFlag{
			Name:  "target-branch",
			Usage: "GitLab target branch",
			//EnvVars: []string{"OA_GITLAB_TARGET_BRANCH"},
		},
		//&cli.BoolFlag{
		//	Name:    "auto-mr",
		//	Usage:   "Automatically create a merge request in GitLab",
		//	EnvVars: []string{"OA_GITLAB_AUTO_MR"},
		//},
	)
}
