package repository

import (
	"fmt"
	"strconv"

	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	model "github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CreateWithRElR(nodeSource interface{}, relation interface{}, nodeTarget interface{}) ([]string, string) {
	var c []string
	var C string
	var excep string

	nSource, properSource := getProperties(nodeSource)
	rel, properRel := getRelationshipConfig(relation)
	nTarget, properTarget := getProperties(nodeTarget)

	// fmt.Println(nSource)
	// fmt.Println(properSource)
	// fmt.Println(rel)
	// fmt.Println(properRel)
	// fmt.Println(nTarget)
	// fmt.Println(properTarget)

	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (t {"+properTarget+"}) WITH t "+
				"MERGE (s{"+properSource+"}) WITH s, t "+
				"CALL apoc.create.addLabels(t, [$node2])YIELD node "+
				"CALL apoc.create.addLabels(s, [$node1]) YIELD node AS N "+
				"WITH s, t, N "+
				"CALL apoc.create.relationship(s, $rel,{"+properRel+"}, t)YIELD rel "+
				"WITH s, t, N "+
				"CALL apoc.path.subgraphAll(N, {maxLevel:0}) YIELD nodes WITH [N in nodes | N {.*, label:labels(N)[0]}] as nodes "+
				"RETURN apoc.convert.toJson(nodes[0]) ",
			map[string]interface{}{"node1": nSource, "node2": nTarget, "rel": rel})
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
		fmt.Println(err)
		return c, err.Error()
	}
	return c, excep
}

func SaveR(title string, released string, m *model.Movie) (model.Movie, string) {
	var movie model.Movie
	mReleased := strconv.FormatInt(m.Released, 10)
	excep := ""
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie {title:$title, released:toInteger($released)}) SET m.title = $mtitle, m.tagline = $mtagline, m.released = toInteger($mreleased) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l",
			map[string]interface{}{"title": title, "released": released, "mtitle": m.Title, "mtagline": m.Tagline, "mreleased": mReleased})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			movie.Title = result.Record().GetByIndex(0).(string)
			movie.Tagline = result.Record().GetByIndex(1).(string)
			movie.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			return nil, nil
		}
		excep = "não encontrado"
		return nil, result.Err()
	})
	if err != nil {
		return movie, "error"
	}
	return movie, excep
}

func DeleteR(title string, released string) string {
	excep := ""
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie {title:$title,released:toInteger($released)}) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l",
			map[string]interface{}{"title": title, "released": released})
		if err != nil {
			return nil, err
		}
		if !result.Next() {
			excep = "não encontrado"
			return nil, nil
		}
		result, err = transaction.Run(
			"MATCH (m:Movie {title:$title,released:toInteger($released)}) DETACH DELETE m",
			map[string]interface{}{"title": title, "released": released})
		if err != nil {
			return nil, err
		}
		return "", nil
	})
	if err != nil {
		excep = "error"
	}
	return excep
}
