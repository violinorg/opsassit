package tests

import (
	"os"
	"testing"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/cmd"
)

func TestAutoMrCmd(t *testing.T) {
	// Set up environment variables
	os.Setenv("OA_GITLAB_AUTOMR_FILE_PATH", "path/to/file")
	os.Setenv("OA_GITLAB_URL", "https://gitlab.example.com")
	os.Setenv("OA_GITLAB_TOKEN", "your-token")
	os.Setenv("OA_GITLAB_PROJECT_ID", "12345")
	os.Setenv("OA_GITLAB_BASE_BRANCH", "main")
	os.Setenv("OA_GITLAB_NEW_BRANCH", "feature-branch")
	os.Setenv("OA_GITLAB_TARGET_BRANCH", "develop")

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.GitlabCmd(),
	}

	// Run the CLI app with the auto-mr command
	err := app.Run([]string{"app", "gitlab", "auto-mr"})
	if err != nil {
		t.Fatalf("Failed to run auto-mr command: %v", err)
	}
}
