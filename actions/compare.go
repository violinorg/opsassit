package actions

import (
	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func LoadVariablesFromYAML(filePath string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CompareKeys(vars1, vars2 map[string]interface{}) ([]string, []string) {
	var onlyInFile1, onlyInFile2 []string

	for key := range vars1 {
		if _, exists := vars2[key]; !exists {
			onlyInFile1 = append(onlyInFile1, key)
		}
	}

	for key := range vars2 {
		if _, exists := vars1[key]; !exists {
			onlyInFile2 = append(onlyInFile2, key)
		}
	}

	return onlyInFile1, onlyInFile2
}

func CompareValues(vars1, vars2 map[string]interface{}) map[string][2]interface{} {
	differences := make(map[string][2]interface{})
	for key, val1 := range vars1 {
		if val2, exists := vars2[key]; exists && !cmp.Equal(val1, val2) {
			differences[key] = [2]interface{}{val1, val2}
		}
	}
	return differences
}

func SaveComparisonResult(resultFilePath string, onlyInFile1, onlyInFile2 []string) error {
	result := struct {
		OnlyInFile1 []string `yaml:"only_in_file1"`
		OnlyInFile2 []string `yaml:"only_in_file2"`
	}{
		OnlyInFile1: onlyInFile1,
		OnlyInFile2: onlyInFile2,
	}

	data, err := yaml.Marshal(result)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(resultFilePath, data, 0644)
}

func saveValuesComparisonResult(resultFilePath string, differences map[string][2]interface{}) error {
	result := struct {
		Differences map[string][2]interface{} `yaml:"differences"`
	}{
		Differences: differences,
	}

	data, err := yaml.Marshal(result)
	if err != nil {
		return err
	}

	return os.WriteFile(resultFilePath, data, 0644)
}
