package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestYaml(t *testing.T) {
	t.Run("TestDiffCmd", TestDiffCmd)
	t.Run("TestStringDiffCmd", TestStringDiffCmd)
}

func TestGitlab(t *testing.T) {
	t.Run("TestGitLabFlags", TestGitLabFlags)
	t.Run("TestAutoMrCmd", TestAutoMrCmd)
}
