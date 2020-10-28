package utils

import (
	"reflect"
	"strings"
)

func Label_Properties_Relationship(rel interface{}) (string, string) {
	var properties, relation string
	fields := reflect.TypeOf(rel)
	values := reflect.ValueOf(rel)
	num := fields.NumField()
	relation = values.Field(0).String()
	for i := 1; i < num; i++ {
		field := strings.ToLower(fields.Field(i).Name)
		value := values.Field(i).Interface()
		properties = properties + field + ": " + TransToString(value)
		if i <= num-2 {
			properties = properties + ", "
		}
	}
	return relation, properties
}
