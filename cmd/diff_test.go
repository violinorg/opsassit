package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDiffCmd(t *testing.T) {
	// Paths to test files
	file1Path := filepath.Join("../tests", "diff", "file1.yaml")
	file2Path := filepath.Join("../tests", "diff", "file2.yaml")
	expectedPath := filepath.Join("../tests", "diff", "expected.yaml")
	expectedPreviewPath := filepath.Join("../tests", "diff", "expected_preview.txt")
	outputPath := filepath.Join("../tests", "diff", "output.yaml")

	// Read expected outputs
	expected, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	expectedPreview, err := os.ReadFile(expectedPreviewPath)
	if err != nil {
		t.Fatalf("Failed to read expected preview output: %v", err)
	}

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		diffCmd(),
	}

	// Set up the environment variables
	os.Setenv("FILE1_PATH", file1Path)
	os.Setenv("FILE2_PATH", file2Path)
	os.Setenv("OA_DRAIN_OUTPUT", outputPath)

	// Run the CLI app with the diff command without --approved
	var buf bytes.Buffer
	app.Writer = &buf
	err = app.Run([]string{"app", "yaml", "diff"})
	if err != nil {
		t.Fatalf("Failed to run diff command: %v", err)
	}

	// Check the preview output
	previewStr := strings.TrimSpace(string(expectedPreview))
	actualPreviewStr := strings.TrimSpace(buf.String())

	if previewStr != actualPreviewStr {
		t.Fatalf("Expected preview:\n%s\nGot:\n%s", previewStr, actualPreviewStr)
	}

	// Reset buffer and run the CLI app with the diff command with --approved
	buf.Reset()
	err = app.Run([]string{"app", "yaml", "diff", "--approved"})
	if err != nil {
		t.Fatalf("Failed to run diff command with --approved: %v", err)
	}

	// Check the output file
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedStr := strings.TrimSpace(string(expected))
	contentStr := strings.TrimSpace(string(content))

	if expectedStr != contentStr {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expectedStr, contentStr)
	}

	// Check the console output
	expectedSuccessMessage := "Successfully drained keys from file2 to output file."
	if !strings.Contains(buf.String(), expectedSuccessMessage) {
		t.Fatalf("Expected success message, got: %s", buf.String())
	}

	// Run the CLI app again to check for "The file is already tuned."
	buf.Reset()
	err = app.Run([]string{"app", "yaml", "diff"})
	if err != nil {
		t.Fatalf("Failed to run diff command: %v", err)
	}
	expectedAlreadyTunedMessage := "The file is already tuned."
	if !strings.Contains(buf.String(), expectedAlreadyTunedMessage) {
		t.Fatalf("Expected 'The file is already tuned.' message, got: %s", buf.String())
	}
}
