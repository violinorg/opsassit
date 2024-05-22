package actions

//
//import (
//	"fmt"
//	"io/ioutil"
//	"strings"
//
//	"github.com/google/go-cmp/cmp"
//	"gopkg.in/yaml.v2"
//)
//
////type OrderedMap struct {
////	Keys   []string
////	Values map[string]interface{}
////}
//
//func LoadVariablesFromYAMLWithOrder(filePath string) (*OrderedMap, error) {
//	content, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		return nil, err
//	}
//
//	var data map[string]interface{}
//	err = yaml.Unmarshal(content, &data)
//	if err != nil {
//		return nil, err
//	}
//
//	orderedMap := &OrderedMap{
//		Keys:   make([]string, 0, len(data)),
//		Values: make(map[string]interface{}),
//	}
//
//	for key, value := range data {
//		orderedMap.Keys = append(orderedMap.Keys, key)
//		orderedMap.Values[key] = value
//	}
//
//	return orderedMap, nil
//}
//
//func GenerateUpdatedYAML(vars1, vars2 *OrderedMap) (string, error) {
//	var builder strings.Builder
//	builder.WriteString("---\n")
//
//	// Preserve the order of keys from vars1
//	for _, key := range vars1.Keys {
//		val1 := vars1.Values[key]
//		if val2, exists := vars2.Values[key]; exists && !cmp.Equal(val1, val2) {
//			builder.WriteString(fmt.Sprintf("# from file2 - %s: %v\n", key, val2))
//		}
//		builder.WriteString(fmt.Sprintf("%s: %v\n", key, val1))
//	}
//
//	// Add keys from vars2 that do not exist in vars1, preserving the order from vars2
//	for _, key := range vars2.Keys {
//		if _, exists := vars1.Values[key]; !exists {
//			builder.WriteString(fmt.Sprintf("%s: %v\n", key, vars2.Values[key]))
//		}
//	}
//
//	return builder.String(), nil
//}
