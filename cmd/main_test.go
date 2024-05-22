package cmd

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("TestDrainCmd", TestDrainCmd)
	t.Run("TestDiffKeysCommand", TestDiffKeysCommand)
	t.Run("TestDiffValuesCommand", TestDiffValuesCommand)
}
