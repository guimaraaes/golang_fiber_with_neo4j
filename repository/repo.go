package repository

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	model "github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
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

func FindR(model interface{}, info map[string]interface{}) ([]string, string) {
	var c []string
	var C string
	node := reflect.TypeOf(model).Elem().Name()
	excep := ""
	var query string
	var m map[string]interface{} = nil
	properties := ""
	i := 0
	for pro, value := range info {
		i = i + 1
		properties = properties + pro + ": " + TransToString(value)
		if i <= len(info)-1 {
			properties = properties + ", "

		}
	}

	// fmt.Println(properties)
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		if info == nil {
			query = "MATCH (m) WHERE labels(m) = [$node] CALL apoc.path.subgraphAll(m, {maxLevel:0}) YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes RETURN apoc.convert.toJson(nodes[0])"
			m = map[string]interface{}{"node": node}
		} else {
			query = "MATCH (m {" + properties + "  }) WHERE labels(m) = [$node] CALL apoc.path.subgraphAll(m, {maxLevel:0}) YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes RETURN apoc.convert.toJson(nodes[0])"
			m = map[string]interface{}{"node": node}
		}
		result, err := transaction.Run(query, m)
		if err != nil {
			return nil, err
		}
		i := 0
		for result.Next() {
			C = result.Record().GetByIndex(0).(string)
			i = i + 1
			c = append(c, C)
		}
		if c == nil {
			excep = "não encontrado"
		}
		return nil, result.Err()
	})
	if err != nil {
		return c, "error"
	}
	return c, excep
}

func CreateR(model interface{}, info map[string]interface{}) ([]string, string) {
	var c []string
	var C string
	// node := reflect.TypeOf(model).Elem().Name()
	excep := ""

	properties := ""
	i := 0
	for pro, value := range info {
		i = i + 1
		properties = properties + pro + ": " + TransToString(value)
		if i <= len(info)-1 {
			properties = properties + ", "

		}
	}
	fmt.Println(properties)
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (mExist {"+properties+"}) WHERE labels(mExist) = [$n] RETURN apoc.convert.toJson(mExist.title) as t",
			map[string]interface{}{"n": "M"})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			excep = "já existe"
			return nil, nil
		}
		result, err = transaction.Run(
			"MERGE (m {"+properties+"}) CALL apoc.path.subgraphAll(m, {maxLevel:0}) YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes RETURN apoc.convert.toJson(nodes[0])",
			nil)
		if err != nil {
			return nil, err
		}
		i := 0
		for result.Next() {
			C = result.Record().GetByIndex(0).(string)
			i = i + 1
			c = append(c, C)
		}
		if c == nil {
			excep = "não encontrado"
		}
		return nil, result.Err()
	})
	if err != nil {
		fmt.Println(err)
		return c, "error"

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
