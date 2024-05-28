package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/cmd"
)

func TestAutoMrCmd(t *testing.T) {
	// Set up environment variables
	srcFilePath := filepath.Join("gitlab", "data", "output_format_all.yaml")
	filePath := filepath.Join("configs", "app_config.yaml")
	os.Setenv("OA_GITLAB_SOURCE_FILE_PATH", srcFilePath)
	os.Setenv("OA_GITLAB_FILE_PATH", filePath)
	os.Setenv("OA_GITLAB_URL", "https://gitlab.com")
	os.Setenv("OA_GITLAB_TOKEN", "glpat-c8jsUhzPpQuic-abxXMX")
	os.Setenv("OA_GITLAB_PROJECT_ID", "58164058")
	os.Setenv("OA_GITLAB_MR_BASE_BRANCH", "main")
	os.Setenv("OA_GITLAB_MR_NEW_BRANCH", "feature/oa-branch")
	os.Setenv("OA_GITLAB_MR_TARGET_BRANCH", "main")

	// Ensure the test directory and file exist
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	err = os.WriteFile(filePath, []byte("test_key: test_value"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.RemoveAll(filepath.Dir(filePath))

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.GitlabCmd(),
	}

	// Run the CLI app with the auto-mr command
	err = app.Run([]string{"app", "gitlab", "auto-mr"})
	if err != nil {
		t.Fatalf("Failed to run auto-mr command: %v", err)
	}
}
