package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func CreateR(m interface{}) ([]string, string) {
	var c []string
	// var C string
	excep := ""

	node, _ := getProperties(m)
	fmt.Println(node)
	properties := map[string]interface{}{}

	// fmt.Println((node))
	fields := reflect.TypeOf(m)
	values := reflect.ValueOf(m)
	num := 2
	for i := 0; i < num; i++ {
		field := strings.ToLower(fields.Field(i).Name)
		value := values.Field(i).Interface()
		// fmt.Println(value)
		properties[field] = (value)
		// if i <= num-2 {
		// 	properties = properties + ", "
		// }
	}
	// properties["field"] = ("value")

	fmt.Println(properties)
	// var p []model.Person
	// var p reflect.TypeOf(m)
	c, _ = FindR(m, properties)

	if c != nil {
		excep = "já existe"
		return nil, excep
	}
	// _, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
	// 	// result, err := transaction.Run(
	// 	// 	"MATCH (mExist {"+properties+"}) "+
	// 	// 		"WHERE $node in labels(mExist) "+
	// 	// 		"CALL apoc.path.subgraphAll(mExist, {maxLevel:0}) YIELD nodes "+
	// 	// 		"WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes "+
	// 	// 		"RETURN apoc.convert.toJson(nodes[0])",
	// 	// 	map[string]interface{}{"node": node})
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }
	// 	// if result.Next() {
	// 	// 	excep = "já existe"
	// 	// 	return nil, nil
	// 	// }

	// 	result, err := transaction.Run(
	// 		"MERGE (m {"+properties+"}) WITH m "+
	// 			"CALL apoc.create.addLabels(m, [$node]) YIELD node "+
	// 			"CALL apoc.path.subgraphAll(node, {maxLevel:0}) "+
	// 			"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes "+
	// 			"RETURN apoc.convert.toJson(nodes[0])",
	// 		map[string]interface{}{"node": node})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	for result.Next() {
	// 		C = result.Record().GetByIndex(0).(string)
	// 		c = append(c, C)
	// 	}
	// 	if c == nil {
	// 		excep = "não encontrado"
	// 	}
	// 	return nil, result.Err()
	// })
	// if err != nil {
	// 	return c, err.Error()
	// }
	return c, excep
}
