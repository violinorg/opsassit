package actions

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v2"
)

func LoadVariablesFromYAML(filePath string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var variables map[string]interface{}
	err = yaml.Unmarshal(data, &variables)
	if err != nil {
		return nil, err
	}

	return variables, nil
}

func LoadVariablesFromYAMLWithOrder(filePath string) (map[string]interface{}, []string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var rawVariables yaml.MapSlice
	err = yaml.Unmarshal(data, &rawVariables)
	if err != nil {
		return nil, nil, err
	}

	variables := make(map[string]interface{})
	var order []string
	for _, item := range rawVariables {
		key := item.Key.(string)
		variables[key] = item.Value
		order = append(order, key)
	}

	return variables, order, nil
}

func CompareValues(vars1, vars2 map[string]interface{}) map[string][2]interface{} {
	allKeys := make(map[string]struct{})
	for k := range vars1 {
		allKeys[k] = struct{}{}
	}
	for k := range vars2 {
		allKeys[k] = struct{}{}
	}

	differences := make(map[string][2]interface{})
	for key := range allKeys {
		val1, ok1 := vars1[key]
		val2, ok2 := vars2[key]
		if !ok1 {
			val1 = nil
		}
		if !ok2 {
			val2 = nil
		}
		if !reflect.DeepEqual(val1, val2) {
			differences[key] = [2]interface{}{val1, val2}
		}
	}

	return differences
}

func CompareKeys(vars1, vars2 map[string]interface{}) (onlyInFile1, onlyInFile2 []string) {
	keysInFile1 := make(map[string]struct{})
	keysInFile2 := make(map[string]struct{})

	for key := range vars1 {
		keysInFile1[key] = struct{}{}
	}
	for key := range vars2 {
		keysInFile2[key] = struct{}{}
	}

	for key := range keysInFile1 {
		if _, exists := keysInFile2[key]; !exists {
			onlyInFile1 = append(onlyInFile1, key)
		}
	}

	for key := range keysInFile2 {
		if _, exists := keysInFile1[key]; !exists {
			onlyInFile2 = append(onlyInFile2, key)
		}
	}

	return
}

func SaveComparisonResult(filePath string, keysOnlyInFile1, keysOnlyInFile2 []string) error {
	content := "Keys only in file1:\n"
	for _, key := range keysOnlyInFile1 {
		content += fmt.Sprintf("- %s\n", key)
	}
	content += "\nKeys only in file2:\n"
	for _, key := range keysOnlyInFile2 {
		content += fmt.Sprintf("- %s\n", key)
	}

	return ioutil.WriteFile(filePath, []byte(content), 0644)
}
