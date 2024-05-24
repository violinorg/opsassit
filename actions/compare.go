package actions

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type OrderedMap struct {
	Keys   []string
	Values map[string]interface{}
}

func LoadVariablesFromYAMLWithOrder(filePath string) (*OrderedMap, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	orderedMap := &OrderedMap{
		Keys:   make([]string, 0, len(data)),
		Values: make(map[string]interface{}),
	}

	for key, value := range data {
		orderedMap.Keys = append(orderedMap.Keys, key)
		orderedMap.Values[key] = value
	}

	sort.Strings(orderedMap.Keys)

	return orderedMap, nil
}

func GenerateUpdatedYAML(vars1, vars2 *OrderedMap) (string, error) {
	var builder strings.Builder

	// Preserve the order of keys from vars1
	for _, key := range vars1.Keys {
		val1 := vars1.Values[key]
		// Enhanced output comparing different values for the same keys
		// if val2, exists := vars2.Values[key]; exists && !cmp.Equal(val1, val2) {
		//	 builder.WriteString(color.New(color.FgHiGreen).Sprintf("# Added from file2 - %s: %v\n", key, val2))
		// }
		builder.WriteString(fmt.Sprintf("%s: %v\n", key, val1))
	}

	// Add keys from vars2 that do not exist in vars1, preserving the order from vars2
	for _, key := range vars2.Keys {
		if _, exists := vars1.Values[key]; !exists {
			builder.WriteString(color.New(color.FgHiGreen).Sprintf("# Added from file2\n%s: %v\n", key, vars2.Values[key]))
		}
	}

	return builder.String(), nil
}

func GenerateValuesComparisonYAML(differences map[string][2]interface{}) (string, error) {
	var builder strings.Builder
	for key, vals := range differences {
		builder.WriteString(fmt.Sprintf("%s: %v -> %v\n", color.New(color.FgHiBlue).Sprint(key), vals[0], color.New(color.FgHiGreen).Sprint(vals[1])))
	}
	builder.WriteString("Comparison completed successfully.\n")
	return builder.String(), nil
}

func GenerateKeysComparisonYAML(onlyInFile1, onlyInFile2 []string) (string, error) {
	var builder strings.Builder
	builder.WriteString(color.New(color.FgHiBlue).Sprintf("\nKeys only in file1:\n"))
	for _, key := range onlyInFile1 {
		builder.WriteString(fmt.Sprintf("- %s\n", key))
	}
	builder.WriteString(color.New(color.FgHiBlue).Sprintf("\nKeys only in file2:\n"))
	for _, key := range onlyInFile2 {
		builder.WriteString(fmt.Sprintf("- %s\n", key))
	}
	builder.WriteString("Comparison completed successfully.\n")
	return builder.String(), nil
}

func CompareKeys(vars1, vars2 *OrderedMap) ([]string, []string) {
	var onlyInFile1, onlyInFile2 []string

	for _, key := range vars1.Keys {
		if _, exists := vars2.Values[key]; !exists {
			onlyInFile1 = append(onlyInFile1, key)
		}
	}

	for _, key := range vars2.Keys {
		if _, exists := vars1.Values[key]; !exists {
			onlyInFile2 = append(onlyInFile2, key)
		}
	}

	return onlyInFile1, onlyInFile2
}

func CompareValues(vars1, vars2 *OrderedMap) map[string][2]interface{} {
	differences := make(map[string][2]interface{})
	for _, key := range vars1.Keys {
		val1 := vars1.Values[key]
		if val2, exists := vars2.Values[key]; exists && !cmp.Equal(val1, val2) {
			differences[key] = [2]interface{}{val1, val2}
		}
	}

	for _, key := range vars2.Keys {
		if _, exists := vars1.Values[key]; exists && !cmp.Equal(vars1.Values[key], vars2.Values[key]) {
			differences[key] = [2]interface{}{vars1.Values[key], vars2.Values[key]}
		}
	}

	return differences
}

func CompareYAMLFilesLineByLine(file1Path, file2Path string) (string, error) {
	file1, err := os.Open(file1Path)
	if err != nil {
		return "", fmt.Errorf("error opening file1: %v", err)
	}
	defer file1.Close()

	file2, err := os.Open(file2Path)
	if err != nil {
		return "", fmt.Errorf("error opening file2: %v", err)
	}
	defer file2.Close()

	var result strings.Builder

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	lineNum := 1
	for scanner1.Scan() || scanner2.Scan() {
		line1 := ""
		if scanner1.Scan() {
			line1 = scanner1.Text()
		}
		line2 := ""
		if scanner2.Scan() {
			line2 = scanner2.Text()
		}

		if line1 != line2 {
			if line1 != "" {
				color.New(color.FgHiRed).Fprintf(&result, "%d: - %s\n", lineNum, line1)
			}
			if line2 != "" {
				color.New(color.FgHiGreen).Fprintf(&result, "%d: + %s\n", lineNum, line2)
			}
		} else {
			result.WriteString(fmt.Sprintf("%d:   %s\n", lineNum, line1))
		}

		lineNum++
	}

	if err := scanner1.Err(); err != nil {
		return "", fmt.Errorf("error reading file1: %v", err)
	}
	if err := scanner2.Err(); err != nil {
		return "", fmt.Errorf("error reading file2: %v", err)
	}

	return result.String(), nil
}
