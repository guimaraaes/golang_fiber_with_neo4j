package utils

import (
	"reflect"
	"strings"
)

func Label_and_Properties(model interface{}) (string, string) {
	properties_model := map[string]interface{}{}
	fields_model := reflect.TypeOf(model)
	values_model := reflect.ValueOf(model)
	num := fields_model.NumField()
	for i := 0; i < num; i++ {
		check := values_model.Field(i).IsZero()
		if !check {
			field := strings.ToLower(fields_model.Field(i).Name)
			value := values_model.Field(i).Interface()
			properties_model[field] = (value)
		}
	}

	node := fields_model.Name()
	properties := ""
	i := 0
	for pro, value := range properties_model {
		i = i + 1
		properties = properties + pro + ": " + TransToString(value)
		if i <= len(properties_model)-1 {
			properties = properties + ", "
		}
	}
	return node, properties
}
