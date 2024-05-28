package tests

import (
	"github.com/violinorg/opsassit/cmd"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGitLabFlags(t *testing.T) {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.GitlabCmd(),
	}

	tests := []struct {
		name     string
		args     []string
		envVars  map[string]string
		expected map[string]string
	}{
		{
			name: "all flags",
			args: []string{
				"app", "gitlab", "auto-mr",
				"--gitlab-url=https://gitlab.com",
				"--gitlab-token=glpat-c8jsUhzPpQuic-abxXMX",
				"--gitlab-project-id=58164058",
				"--mr-base-branch=main",
				"--mr-new-branch=feature/oa-branch",
				"--mr-target-branch=main",
				"--mr-title=Example MR",
				"--mr-description=Gitlab flags test",
				"--mr-squash=true",
				"--mr-remove-source-branch=true",
				"--src-file-path=gitlab/data/output_format_all.yaml",
				"--file-path=configs/app_config.yaml",
			},
			envVars: map[string]string{},
			expected: map[string]string{
				"gitlab-url":              "https://gitlab.com",
				"gitlab-token":            "glpat-c8jsUhzPpQuic-abxXMX",
				"gitlab-project-id":       "58164058",
				"mr-base-branch":          "main",
				"mr-new-branch":           "feature/oa-branch",
				"mr-target-branch":        "main",
				"mr-title":                "Example MR",
				"mr-description":          "Gitlab flags test",
				"mr-squash":               "true",
				"mr-remove-source-branch": "true",
				"src-file-path":           "gitlab/data/output_format_all.yaml",
				"file-path":               "configs/app_config.yaml",
			},
		},
		{
			name: "env vars",
			args: []string{"app", "gitlab", "auto-mr"},
			envVars: map[string]string{
				"OA_GITLAB_URL":                     "https://gitlab.com",
				"OA_GITLAB_TOKEN":                   "glpat-c8jsUhzPpQuic-abxXMX",
				"OA_GITLAB_PROJECT_ID":              "58164058",
				"OA_GITLAB_MR_BASE_BRANCH":          "main",
				"OA_GITLAB_MR_NEW_BRANCH":           "feature/oa-branch",
				"OA_GITLAB_MR_TARGET_BRANCH":        "main",
				"OA_GITLAB_MR_TITLE":                "Example MR",
				"OA_GITLAB_MR_DESCRIPTION":          "Gitlab flags test",
				"OA_GITLAB_MR_SQUASH":               "true",
				"OA_GITLAB_MR_REMOVE_SOURCE_BRANCH": "true",
				"OA_GITLAB_SRC_FILE_PATH":           "gitlab/data/output_format_all.yaml",
				"OA_GITLAB_FILE_PATH":               "configs/app_config.yaml",
			},
			expected: map[string]string{
				"gitlab-url":              "https://gitlab.com",
				"gitlab-token":            "glpat-c8jsUhzPpQuic-abxXMX",
				"gitlab-project-id":       "58164058",
				"mr-base-branch":          "main",
				"mr-new-branch":           "feature/oa-branch",
				"mr-target-branch":        "main",
				"mr-title":                "Example MR",
				"mr-description":          "Gitlab flags test",
				"mr-squash":               "true",
				"mr-remove-source-branch": "true",
				"src-file-path":           "gitlab/data/output_format_all.yaml",
				"file-path":               "configs/app_config.yaml",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.envVars {
				os.Setenv(key, value)
			}

			setFlags := map[string]string{}
			app.Commands[0].Subcommands[0].Action = func(c *cli.Context) error {
				setFlags["gitlab-url"] = c.String("gitlab-url")
				setFlags["gitlab-token"] = c.String("gitlab-token")
				setFlags["gitlab-project-id"] = c.String("gitlab-project-id")
				setFlags["mr-base-branch"] = c.String("mr-base-branch")
				setFlags["mr-new-branch"] = c.String("mr-new-branch")
				setFlags["mr-target-branch"] = c.String("mr-target-branch")
				setFlags["mr-title"] = c.String("mr-title")
				setFlags["mr-description"] = c.String("mr-description")
				setFlags["mr-squash"] = c.String("mr-squash")
				setFlags["mr-remove-source-branch"] = c.String("mr-remove-source-branch")
				setFlags["src-file-path"] = c.String("src-file-path")
				setFlags["file-path"] = c.String("file-path")
				return nil
			}

			err := app.Run(test.args)
			assert.NoError(t, err)

			for key, expectedValue := range test.expected {
				assert.Equal(t, expectedValue, setFlags[key])
			}

			for key := range test.envVars {
				os.Unsetenv(key)
			}
		})
	}
}
