package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/cmd"
)

func TestDiffCmd(t *testing.T) {
	// Paths to test files
	file1Path := filepath.Join("diff", "file1.yaml")
	file2Path := filepath.Join("diff", "file2.yaml")
	expectedPath := filepath.Join("diff", "expected.yaml")
	outputPath := filepath.Join("diff", "output.yaml")

	// Read expected outputs
	expected, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.YamlCmd(),
	}

	// Set up the environment variables
	os.Setenv("FILE1_PATH", file1Path)
	os.Setenv("FILE2_PATH", file2Path)
	os.Setenv("OA_DRAIN_OUTPUT", outputPath)

	// Run the CLI app with the diff command without --approved
	err = app.Run([]string{"app", "yaml", "diff", file1Path, file2Path})
	if err != nil {
		t.Fatalf("Failed to run diff command: %v", err)
	}

	// Check the preview output
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedStr := strings.TrimSpace(string(expected))
	contentStr := strings.TrimSpace(string(content))

	assert.Equal(t, expectedStr, contentStr, "The content should match the expected output")

	// Check the console output
	expectedSuccessMessage := "Successfully drained keys from file2 to output file."
	if !strings.Contains(contentStr, expectedSuccessMessage) {
		t.Fatalf("Expected success message, got: %s", contentStr)
	}

	// Run the CLI app again to check for "The file is already tuned."
	err = app.Run([]string{"app", "yaml", "diff", file1Path, file2Path})
	if err != nil {
		t.Fatalf("Failed to run diff command: %v", err)
	}
	expectedAlreadyTunedMessage := "The file is already tuned."
	if !strings.Contains(contentStr, expectedAlreadyTunedMessage) {
		t.Fatalf("Expected 'The file is already tuned.' message, got: %s", contentStr)
	}
}

func TestStringDiffCmd(t *testing.T) {
	// Paths to test files
	file1Path := filepath.Join("diff", "file1.yaml")
	file2Path := filepath.Join("diff", "file2.yaml")
	expectedStringDiffPath := filepath.Join("diff", "expected_string_diff.txt")
	outputStringDiffPath := filepath.Join("diff", "output_string_diff.txt")

	// Read expected outputs
	expectedStringDiff, err := os.ReadFile(expectedStringDiffPath)
	if err != nil {
		t.Fatalf("Failed to read expected string diff output: %v", err)
	}

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.YamlCmd(),
	}

	// Set up the environment variables
	os.Setenv("FILE1_PATH", file1Path)
	os.Setenv("FILE2_PATH", file2Path)
	os.Setenv("OA_YAML_DIFF_OUTPUT_PATH", outputStringDiffPath)

	// Run the CLI app with the diff command with --format string
	err = app.Run([]string{"app", "yaml", "diff", "--format", "string", file1Path, file2Path})
	if err != nil {
		t.Fatalf("Failed to run diff command with --format string: %v", err)
	}

	// Check the console output
	content, err := os.ReadFile(outputStringDiffPath) // Replace with actual output capture
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedStr := strings.TrimSpace(string(expectedStringDiff))
	contentStr := strings.TrimSpace(string(content))

	assert.Equal(t, expectedStr, contentStr, "The content should match the expected string diff output")
}
