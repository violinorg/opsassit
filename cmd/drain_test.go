package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDrainCmd(t *testing.T) {
	// Create a temporary directory for testing
	dir, err := os.MkdirTemp("", "drain_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	// Create test files
	file1Path := dir + "/file1.yaml"
	file2Path := dir + "/file2.yaml"

	err = os.WriteFile(file1Path, []byte(`var1: 10
var2: 20
var3: 30
`), 0644)
	if err != nil {
		t.Fatalf("Failed to write to file1: %v", err)
	}

	err = os.WriteFile(file2Path, []byte(`var2: 25
var3: 30
var4: 40
`), 0644)
	if err != nil {
		t.Fatalf("Failed to write to file2: %v", err)
	}

	// Set up the CLI app
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		drainCmd(),
	}

	// Set up the environment variables
	os.Setenv("FILE1_PATH", file1Path)
	os.Setenv("FILE2_PATH", file2Path)

	// Run the CLI app with the drain command
	var buf bytes.Buffer
	app.Writer = &buf
	err = app.Run([]string{"app", "drain"})
	if err != nil {
		t.Fatalf("Failed to run drain command: %v", err)
	}

	// Check the output
	expected := "---\nvar1: 10\n# from file2 - var2: 25\nvar2: 20\nvar3: 30\nvar4: 40\n\n# OpsAssist Verified\n"
	content, err := os.ReadFile(file1Path)
	if err != nil {
		t.Fatalf("Failed to read file1: %v", err)
	}
	if string(content) != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, string(content))
	}

	// Check the console output
	if !bytes.Contains(buf.Bytes(), []byte("Successfully drained keys from file2 to file1.")) {
		t.Fatalf("Expected success message, got: %s", buf.String())
	}

	// Run the CLI app again to check for "The file is already tuned."
	buf.Reset()
	err = app.Run([]string{"app", "drain"})
	if err != nil {
		t.Fatalf("Failed to run drain command: %v", err)
	}
	if !bytes.Contains(buf.Bytes(), []byte("The file is already tuned.")) {
		t.Fatalf("Expected 'The file is already tuned.' message, got: %s", buf.String())
	}
}
