package tests

//
//import (
//	"os"
//	"path/filepath"
//	"strings"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/urfave/cli/v2"
//	"github.com/violinorg/opsassit/cmd"
//)
//
//func setupTestFiles(t *testing.T, testDir string, file1Content, file2Content, expectedContent, expectedStringDiffContent string) (file1Path, file2Path, expectedPath, expectedStringDiffPath, outputPath string) {
//	// Paths to test files
//	file1Path = filepath.Join(testDir, "file1.yaml")
//	file2Path = filepath.Join(testDir, "file2.yaml")
//	expectedPath = filepath.Join(testDir, "expected.yaml")
//	outputPath = filepath.Join(testDir, "output_test.txt")
//
//	// Ensure the test directory exists
//	err := os.MkdirAll(testDir, 0755)
//	if err != nil {
//		t.Fatalf("Failed to create test directory: %v", err)
//	}
//
//	// Write test files
//	err = os.WriteFile(file1Path, []byte(file1Content), 0644)
//	if err != nil {
//		t.Fatalf("Failed to write file1: %v", err)
//	}
//
//	err = os.WriteFile(file2Path, []byte(file2Content), 0644)
//	if err != nil {
//		t.Fatalf("Failed to write file2: %v", err)
//	}
//
//	err = os.WriteFile(expectedPath, []byte(expectedContent), 0644)
//	if err != nil {
//		t.Fatalf("Failed to write expected file: %v", err)
//	}
//
//	err = os.WriteFile(expectedStringDiffPath, []byte(expectedStringDiffContent), 0644)
//	if err != nil {
//		t.Fatalf("Failed to write expected string diff file: %v", err)
//	}
//
//	return
//}
//
//func runDiffTest(t *testing.T, app *cli.App, file1Path, file2Path, outputPath, expectedOutput string, args ...string) {
//	// Set up the environment variables
//	os.Setenv("FILE1_PATH", file1Path)
//	os.Setenv("FILE2_PATH", file2Path)
//	os.Setenv("OA_YAML_DIFF_OUTPUT_PATH", outputPath)
//
//	// Run the CLI app with the provided arguments
//	err := app.Run(args)
//	if err != nil {
//		t.Fatalf("Failed to run diff command: %v", err)
//	}
//
//	// Check the preview output
//	content, err := os.ReadFile(outputPath)
//	if err != nil {
//		t.Fatalf("Failed to read output file: %v", err)
//	}
//
//	expectedStr := strings.TrimSpace(expectedOutput)
//	contentStr := strings.TrimSpace(string(content))
//
//	assert.Equal(t, expectedStr, contentStr, "The content should match the expected output")
//}
//
//func TestDiffCmd(t *testing.T) {
//	file1Content := `
//var1: 10
//var2: 20
//var3: 30
//`
//	file2Content := `
//var2: 25
//var3: 30
//var4: 40
//`
//	expectedContent := `
//# OpsAssist Verified
//var1: 10
//# from file2 - var2: 25
//var2: 20
//var3: 30
//var4: 40
//`
//	expectedStringDiffContent := `
//1:   var1: 10
//2: - var2: 20
//2: + var2: 25
//3:   var3: 30
//4: + var4: 40
//`
//
//	file1Path, file2Path, expectedPath, expectedStringDiffPath, outputPath := setupTestFiles(t, "tests/diff", file1Content, file2Content, expectedContent, expectedStringDiffContent)
//
//	// Set up the CLI app
//	app := cli.NewApp()
//	app.Commands = []*cli.Command{
//		cmd.YamlCmd(),
//	}
//
//	// Read expected outputs
//	expected, err := os.ReadFile(expectedPath)
//	if err != nil {
//		t.Fatalf("Failed to read expected output: %v", err)
//	}
//
//	runDiffTest(t, app, file1Path, file2Path, outputPath, string(expected), "app", "yaml", "diff", file1Path, file2Path)
//}
//
//func TestStringDiffCmd(t *testing.T) {
//	file1Content := `
//var1: 10
//var2: 20
//var3: 30
//`
//	file2Content := `
//var2: 25
//var3: 30
//var4: 40
//`
//	expectedContent := `
//# OpsAssist Verified
//var1: 10
//# from file2 - var2: 25
//var2: 20
//var3: 30
//var4: 40
//`
//	expectedStringDiffContent := `
//1:   var1: 10
//2: - var2: 20
//2: + var2: 25
//3:   var3: 30
//4: + var4: 40
//`
//
//	file1Path, file2Path, expectedPath, expectedStringDiffPath, outputPath := setupTestFiles(t, "tests/diff", file1Content, file2Content, expectedContent, expectedStringDiffContent)
//
//	// Set up the CLI app
//	app := cli.NewApp()
//	app.Commands = []*cli.Command{
//		cmd.YamlCmd(),
//	}
//
//	// Read expected outputs
//	expectedStringDiff, err := os.ReadFile(expectedStringDiffPath)
//	if err != nil {
//		t.Fatalf("Failed to read expected string diff output: %v", err)
//	}
//
//	runDiffTest(t, app, file1Path, file2Path, outputPath, string(expectedStringDiff), "app", "yaml", "diff", "--format", "string", file1Path, file2Path)
//}
