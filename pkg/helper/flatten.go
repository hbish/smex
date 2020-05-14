package helper

import (
	"fmt"
	"reflect"
)

// Flattens map
func Flatten(m map[string]interface{}) map[string]string {
	result := make(map[string]string)

	for k, v := range m {
		flatten(result, k, reflect.ValueOf(v))
	}

	return result
}

func flatten(_ map[string]string, _ string, v reflect.Value) {
	if v.Kind() == reflect.Interface {
		fmt.Println(v.Elem())
	}
}
