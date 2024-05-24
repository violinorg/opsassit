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

func runDiffTest(t *testing.T, expectedPath, outputPath string, args ...string) {
	// Paths to test files
	file1Path := filepath.Join("diff/data", "file1.yaml")
	file2Path := filepath.Join("diff/data", "file2.yaml")

	// Set up the environment variables
	os.Setenv("OA_YAML_DIFF_FILE1_PATH", file1Path)
	os.Setenv("OA_YAML_DIFF_FILE2_PATH", file2Path)
	os.Setenv("OA_YAML_DIFF_OUTPUT_PATH", outputPath)

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.YamlCmd(),
	}

	// Read expected outputs
	expectedContent, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read expected string diff output: %v", err)
	}

	err = app.Run(args)
	if err != nil {
		t.Fatalf("Failed to run command with args: %v", err)
	}

	// Check the console output
	content, err := os.ReadFile(outputPath) // Replace with actual output capture
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedStr := strings.TrimSpace(string(expectedContent))
	contentStr := strings.TrimSpace(string(content))

	assert.Equal(t, expectedStr, contentStr, "The content should match the expected string diff output")
}

func TestDiffCmd(t *testing.T) {
	// Paths to test files
	expectedPath := filepath.Join("diff/expected", "expected_format_all.txt")
	outputPath := filepath.Join("diff/output", "output_format_all.txt")

	// Running test
	runDiffTest(t, expectedPath, outputPath, "app", "yaml", "diff", "--approved")
}

func TestKeysDiffCmd(t *testing.T) {
	// Paths to test files
	expectedPath := filepath.Join("diff/expected", "expected_format_keys.txt")
	outputPath := filepath.Join("diff/output", "output_format_keys.txt")

	// Running test
	runDiffTest(t, expectedPath, outputPath, "app", "yaml", "diff", "--approved", "--format", "keys")
}

func TestStringDiffCmd(t *testing.T) {
	// Paths to test files
	expectedPath := filepath.Join("diff/expected", "expected_format_string.txt")
	outputPath := filepath.Join("diff/output", "output_format_string.txt")

	// Running test
	runDiffTest(t, expectedPath, outputPath, "app", "yaml", "diff", "--approved", "--format", "string")
}

func TestValuesDiffCmd(t *testing.T) {
	// Paths to test files
	expectedPath := filepath.Join("diff/expected", "expected_format_values.txt")
	outputPath := filepath.Join("diff/output", "output_format_values.txt")

	// Running test
	runDiffTest(t, expectedPath, outputPath, "app", "yaml", "diff", "--approved", "--format", "values")
}
