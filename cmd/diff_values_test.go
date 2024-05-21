package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDiffValuesCmd(t *testing.T) {
	// Paths to test files
	file1Path := filepath.Join("tests", "diff_values", "file1.yaml")
	file2Path := filepath.Join("tests", "diff_values", "file2.yaml")
	expectedOutputPath := filepath.Join("tests", "diff_values", "expected_output.txt")

	// Read expected output
	expectedOutput, err := os.ReadFile(expectedOutputPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		diffValuesCmd(),
	}

	// Run the CLI app with the diff values command
	var buf bytes.Buffer
	app.Writer = &buf
	err = app.Run([]string{"app", "values", file1Path, file2Path})
	if err != nil {
		t.Fatalf("Failed to run diff values command: %v", err)
	}

	// Trim spaces and newlines for a more lenient comparison
	expectedStr := strings.TrimSpace(string(expectedOutput))
	actualStr := strings.TrimSpace(buf.String())

	if expectedStr != actualStr {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expectedStr, actualStr)
	}
}
