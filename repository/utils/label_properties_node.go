package utils

import (
	"reflect"
	"strings"
)

func Model_Properties_To_Params(model interface{}) map[string]interface{} {
	properties := map[string]interface{}{}
	fields := reflect.TypeOf(model)
	values := reflect.ValueOf(model)
	num := fields.NumField()
	for i := 0; i < num; i++ {
		check := values.Field(i).IsZero()
		if !check {
			field := strings.ToLower(fields.Field(i).Name)
			value := values.Field(i).Interface()
			properties[field] = (value)
		}
	}
	return properties
}

func Label_Properties_Node(model interface{}, info map[string]interface{}) (string, string) {
	node := reflect.TypeOf(model).Name()
	properties := ""
	i := 0
	for pro, value := range info {
		i = i + 1
		properties = properties + pro + ": " + TransToString(value)
		if i <= len(info)-1 {
			properties = properties + ", "
		}
	}
	return node, properties
}
