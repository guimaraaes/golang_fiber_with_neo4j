package utils

import (
	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func QueryCall(query string, m map[string]interface{}) ([]string, string) {
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
