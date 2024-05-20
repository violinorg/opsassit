package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDiffKeysCmd(t *testing.T) {
	// Paths to test files
	file1Path := filepath.Join("tests", "diff_keys", "file1.yaml")
	file2Path := filepath.Join("tests", "diff_keys", "file2.yaml")
	expectedOutputPath := filepath.Join("tests", "diff_keys", "expected_output.txt")

	// Read expected output
	expectedOutput, err := os.ReadFile(expectedOutputPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		diffKeysCmd(),
	}

	// Run the CLI app with the diff keys command
	var buf bytes.Buffer
	app.Writer = &buf
	err = app.Run([]string{"app", "keys", file1Path, file2Path})
	if err != nil {
		t.Fatalf("Failed to run diff keys command: %v", err)
	}

	// Trim spaces and newlines for a more lenient comparison
	expectedStr := strings.TrimSpace(string(expectedOutput))
	actualStr := strings.TrimSpace(buf.String())

	if !strings.Contains(actualStr, expectedStr) {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expectedStr, actualStr)
	}
}
