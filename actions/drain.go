package actions

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func LoadVariablesFromYAMLWithOrder(filePath string) (map[string]interface{}, []string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, nil, err
	}

	var order []string
	for key := range data {
		order = append(order, key)
	}

	return data, order, nil
}

func GenerateUpdatedYAML(vars1, vars2 map[string]interface{}, order1 []string) (string, error) {
	var builder strings.Builder
	builder.WriteString("---\n")

	// Preserve the order of keys from file1
	for _, key := range order1 {
		val1 := vars1[key]
		if val2, exists := vars2[key]; exists && !cmp.Equal(val1, val2) {
			builder.WriteString(fmt.Sprintf("# from file2 - %s: %v\n", key, val2))
		}
		builder.WriteString(fmt.Sprintf("%s: %v\n", key, val1))
	}

	// Add keys from file2 that do not exist in file1
	for key, val2 := range vars2 {
		if _, exists := vars1[key]; !exists {
			builder.WriteString(fmt.Sprintf("%s: %v\n", key, val2))
		}
	}

	return builder.String(), nil
}
