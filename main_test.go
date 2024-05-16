package main

import (
	"reflect"
	"testing"

	"github.com/violinorg/opsassit/actions"
)

func TestLoadVariablesFromYAML(t *testing.T) {
	tests := []struct {
		filePath string
		expected map[string]interface{}
	}{
		{
			filePath: "file1.yaml",
			expected: map[string]interface{}{
				"var1": 10,
				"var2": 20,
				"var3": 30,
			},
		},
		{
			filePath: "file2.yaml",
			expected: map[string]interface{}{
				"var1": 10,
				"var2": 25,
				"var3": 30,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			vars, err := actions.LoadVariablesFromYAML(tt.filePath)
			if err != nil {
				t.Fatalf("Error loading YAML file: %v", err)
			}

			if !reflect.DeepEqual(vars, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, vars)
			}
		})
	}
}

func TestCompareValues(t *testing.T) {
	tests := []struct {
		vars1, vars2 map[string]interface{}
		expected     map[string][2]interface{}
	}{
		{
			vars1: map[string]interface{}{
				"var1": 10,
				"var2": 20,
				"var3": 30,
			},
			vars2: map[string]interface{}{
				"var1": 10,
				"var2": 25,
				"var3": 30,
			},
			expected: map[string][2]interface{}{
				"var2": {20, 25},
			},
		},
		{
			vars1: map[string]interface{}{
				"var1": 10,
				"var2": 20,
			},
			vars2: map[string]interface{}{
				"var1": 10,
				"var2": 20,
				"var3": 30,
			},
			expected: map[string][2]interface{}{
				"var3": {nil, 30},
			},
		},
		{
			vars1: map[string]interface{}{
				"var1": 10,
				"var2": 20,
			},
			vars2: map[string]interface{}{
				"var1": 10,
				"var3": 30,
			},
			expected: map[string][2]interface{}{
				"var2": {20, nil},
				"var3": {nil, 30},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			differences := actions.CompareValues(tt.vars1, tt.vars2)
			if !reflect.DeepEqual(differences, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, differences)
			}
		})
	}
}

func TestCompareKeys(t *testing.T) {
	tests := []struct {
		vars1, vars2                             map[string]interface{}
		expectedOnlyInFile1, expectedOnlyInFile2 []string
	}{
		{
			vars1: map[string]interface{}{
				"var1": 10,
				"var2": 20,
				"var3": 30,
			},
			vars2: map[string]interface{}{
				"var1": 10,
				"var2": 25,
				"var4": 30,
			},
			expectedOnlyInFile1: []string{"var3"},
			expectedOnlyInFile2: []string{"var4"},
		},
		{
			vars1: map[string]interface{}{
				"var1": 10,
				"var2": 20,
			},
			vars2: map[string]interface{}{
				"var1": 10,
				"var2": 20,
				"var3": 30,
			},
			expectedOnlyInFile1: []string{},
			expectedOnlyInFile2: []string{"var3"},
		},
		{
			vars1: map[string]interface{}{
				"var1": 10,
				"var2": 20,
			},
			vars2: map[string]interface{}{
				"var1": 10,
				"var3": 30,
			},
			expectedOnlyInFile1: []string{"var2"},
			expectedOnlyInFile2: []string{"var3"},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			onlyInFile1, onlyInFile2 := actions.CompareKeys(tt.vars1, tt.vars2)
			if !reflect.DeepEqual(onlyInFile1, tt.expectedOnlyInFile1) {
				t.Errorf("Expected only in file1: %v, got %v", tt.expectedOnlyInFile1, onlyInFile1)
			}
			if !reflect.DeepEqual(onlyInFile2, tt.expectedOnlyInFile2) {
				t.Errorf("Expected only in file2: %v, got %v", tt.expectedOnlyInFile2, onlyInFile2)
			}
		})
	}
}
