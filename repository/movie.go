package repository

import (
	"strconv"

	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	"github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func Find(info ...string) ([]model.Movie, string) {
	var movie []model.Movie
	excep := ""
	var query string
	var m map[string]interface{} = nil
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		if info == nil {
			query = "MATCH (m:Movie) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l LIMIT 10"
		} else {
			query = "MATCH (m:Movie {title:$title,released:toInteger($released)}) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l"
			m = map[string]interface{}{"title": info[0], "released": info[1]}
		}
		result, err := transaction.Run(query, m)
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
		if movie == nil {
			excep = "não encontrado"
		}
		return nil, result.Err()
	})
	if err != nil {
		return nil, "error"
	}
	return movie, excep
}

func Create(m *model.Movie) (model.Movie, string) {
	var movie model.Movie
	released := strconv.FormatInt(m.Released, 10)
	excep := ""
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (mExist:N {title:$mtitle, tagline:$mtagline, released:toInteger($mreleased)}) RETURN apoc.convert.toJson(mExist.title) as t, apoc.convert.toJson(mExist.tagline) as tg, apoc.convert.toJson(mExist.released) as l",
			map[string]interface{}{"mtitle": m.Title, "mtagline": m.Tagline, "mreleased": released})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			movie.Title = result.Record().GetByIndex(0).(string)
			movie.Tagline = result.Record().GetByIndex(1).(string)
			movie.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			excep = "já existe"
			return nil, nil
		}
		result, err = transaction.Run(
			"MERGE (m:N {title:$mtitle, tagline:$mtagline, released:toInteger($mreleased)}) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l",
			map[string]interface{}{"mtitle": m.Title, "mtagline": m.Tagline, "mreleased": released})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			movie.Title = result.Record().GetByIndex(0).(string)
			movie.Tagline = result.Record().GetByIndex(1).(string)
			movie.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			return nil, nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return movie, "error"
	}
	return movie, excep

}

func Save(title string, released string, m *model.Movie) (model.Movie, string) {
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

func Delete(title string, released string) string {
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
