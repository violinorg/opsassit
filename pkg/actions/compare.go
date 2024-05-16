package cmd

import (
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

func LoadVariablesFromYAML(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
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
