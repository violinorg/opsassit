//package actions
//
//import (
//	"gopkg.in/yaml.v2"
//	"io/ioutil"
//)
//
//func LoadVariablesFromYAML(filePath string) (map[string]interface{}, error) {
//	data, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		return nil, err
//	}
//
//	var variables map[string]interface{}
//	err = yaml.Unmarshal(data, &variables)
//	if err != nil {
//		return nil, err
//	}
//
//	return variables, nil
//}
//
//func LoadVariablesFromYAMLWithOrder(filePath string) (map[string]interface{}, []string, error) {
//	data, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	var rawVariables yaml.MapSlice
//	err = yaml.Unmarshal(data, &rawVariables)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	variables := make(map[string]interface{})
//	var order []string
//	for _, item := range rawVariables {
//		key := item.Key.(string)
//		variables[key] = item.Value
//		order = append(order, key)
//	}
//
//	return variables, order, nil
//}
