package repository

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func TransToString(data interface{}) (res string) {
	switch v := data.(type) {
	case float64:
		res = strconv.FormatFloat(data.(float64), 'f', 6, 64)
	case float32:
		res = strconv.FormatFloat(float64(data.(float32)), 'f', 6, 32)
	case int:
		res = strconv.FormatInt(int64(data.(int)), 10)
	case int64:
		res = strconv.FormatInt(data.(int64), 10)
	case uint:
		res = strconv.FormatUint(uint64(data.(uint)), 10)
	case uint64:
		res = strconv.FormatUint(data.(uint64), 10)
	case uint32:
		res = strconv.FormatUint(uint64(data.(uint32)), 10)
	case json.Number:
		res = data.(json.Number).String()
	case string:
		res = "'" + data.(string) + "'"
	case []byte:
		res = string(v)
	case bool:
		res = string(strconv.FormatBool(data.(bool)))
	default:
		res = ""
	}
	return
}

func getProperties(model interface{}) (string, string) {
	var properties string
	var node string
	node = reflect.TypeOf(model).Name()
	// fmt.Println((node))
	fields := reflect.TypeOf(model)
	values := reflect.ValueOf(model)
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := strings.ToLower(fields.Field(i).Name)
		value := values.Field(i).Interface()
		// fmt.Println(value)
		properties = properties + field + ": " + TransToString(value)
		if i <= num-2 {
			properties = properties + ", "
		}
	}
	return node, properties
}

func getRelationshipConfig(rel interface{}) (string, string) {
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
