package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/cmd"
)

func TestGitLabAll(t *testing.T) {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.GitlabCmd(),
	}

	for _, test := range GitLabTestCases {
		t.Run(test.Name, func(t *testing.T) {

			for key, value := range test.EnvVars {
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

			err := app.Run(test.Args)
			assert.NoError(t, err)

			for key, expectedValue := range test.Expected {
				assert.Equal(t, expectedValue, setFlags[key])
			}

			for key := range test.EnvVars {
				os.Unsetenv(key)
			}
		})
	}
}
