package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDiffKeysCmd(t *testing.T) {
	// Create a temporary directory for testing
	dir, err := os.MkdirTemp("", "diff_keys_test")
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
		diffKeysCmd(),
	}

	// Run the CLI app with the diff keys command
	var buf bytes.Buffer
	app.Writer = &buf
	err = app.Run([]string{"app", "keys", file1Path, file2Path})
	if err != nil {
		t.Fatalf("Failed to run diff keys command: %v", err)
	}

	// Check the console output
	if !bytes.Contains(buf.Bytes(), []byte("Keys only in file1:")) ||
		!bytes.Contains(buf.Bytes(), []byte("Keys only in file2:")) ||
		!bytes.Contains(buf.Bytes(), []byte("Comparison completed successfully.")) {
		t.Fatalf("Expected keys comparison output, got: %s", buf.String())
	}
}
