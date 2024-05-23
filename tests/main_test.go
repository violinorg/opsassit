package tests

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("TestDiffCmd", TestDiffCmd)
	t.Run("TestAutoMrCmd", TestAutoMrCmd)
}
