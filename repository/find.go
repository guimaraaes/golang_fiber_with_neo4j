package repository

import (
	"fmt"
	"reflect"

	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func FindR(model interface{}, info map[string]interface{}) ([]string, string) {
	var c []string
	var query string
	var m map[string]interface{} = nil

	node := reflect.TypeOf(model).Name()
	fmt.Println(node)
	properties := ""
	i := 0
	for pro, value := range info {
		i = i + 1
		properties = properties + pro + ": " + TransToString(value)
		if i <= len(info)-1 {
			properties = properties + ", "
		}
	}

	if info == nil {
		query = "MATCH (m) WHERE $node in labels(m) " +
			"CALL apoc.path.subgraphAll(m, {maxLevel:0}) " +
			"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes " +
			"RETURN apoc.convert.toJson(nodes[0])"
		m = map[string]interface{}{"node": node}
	} else {
		query = "MATCH (m {" + properties + "  }) WHERE $node in labels(m) " +
			"CALL apoc.path.subgraphAll(m, {maxLevel:0}) " +
			"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes " +
			"RETURN apoc.convert.toJson(nodes[0])"
		m = map[string]interface{}{"node": node}
	}
	c, err := queryCall(query, m)
	return c, err
}

func queryCall(query string, m map[string]interface{}) ([]string, string) {
	var c []string
	var C string
	excep := ""
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(query, m)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			C = result.Record().GetByIndex(0).(string)
			c = append(c, C)
		}
		if c == nil {
			excep = "não encontrado"
		}
		return nil, result.Err()
	})

	if err != nil {
		excep = err.Error()
	}
	return c, excep

}
