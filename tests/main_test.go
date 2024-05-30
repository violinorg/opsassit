package tests

import (
	"testing"
)

func TestYaml(t *testing.T) {
	t.Run("TestDiffCmd", TestDiffCmd)
	t.Run("TestStringDiffCmd", TestStringDiffCmd)
}

func TestGitlab(t *testing.T) {
	t.Run("TestGitLabAll", TestGitLabAll)
}
