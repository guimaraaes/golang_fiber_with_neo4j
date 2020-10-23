package repository

import (
	"strconv"

	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	"github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func All() ([]model.Movie, string) {
	// movies := make([]model.Movie, 0)
	var movie []model.Movie
	excep := ""
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l LIMIT 10",
			nil)
		if err != nil {
			return nil, err
		}

		if !result.Next() {
			excep = "não encontrado"
			return nil, result.Err()
		}

		i := 0
		var m model.Movie
		for result.Next() {
			m.Title = result.Record().GetByIndex(0).(string)
			m.Tagline = result.Record().GetByIndex(1).(string)
			m.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			movie = append(movie, m)
			i = i + 1
		}
		return nil, nil
	})
	if err != nil {
		return nil, "error"
	}

	return movie, excep
}

func Find(info ...string) ([]model.Movie, string) {
	// info[0]
	var movie []model.Movie
	excep := ""
	var query string
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		if info == nil {
			query = "MATCH (m:Movie) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l LIMIT 10"
		} else {
			query = "MATCH (m:Movie {title:$title,released:toInteger($released)}) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l"
		}
		result, err := transaction.Run(query,
			map[string]interface{}{"title": info[0], "released": info[1]})
		if err != nil {
			return nil, err
		}
		i := 0
		var m model.Movie
		for result.Next() {
			m.Title = result.Record().GetByIndex(0).(string)
			m.Tagline = result.Record().GetByIndex(1).(string)
			m.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			movie = append(movie, m)
			i = i + 1
		}
		return nil, nil

		excep = "não encontrado"
		return nil, result.Err()
	})
	if err != nil {
		return movie, "error"
	}
	return movie, excep
}

// func Create() {

// }

// func Save() {

// }

// func Delete() {

// }
