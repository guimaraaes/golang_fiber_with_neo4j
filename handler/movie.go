package handler

import (
	url "net/url"
	strconv "strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	"github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// GetMovie godoc
// @Tags movie
// @Summary Show all movies
// @Produce  json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie [get]
func GetMovie(c *fiber.Ctx) error {
	movies := make([]model.Movie, 10)
	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.released) as l LIMIT 10",
			nil)
		if err != nil {
			return nil, err
		}

		if !result.Next() {
			return nil, result.Err()
		}

		i := 0
		for result.Next() {
			title, _ := result.Record().Get("t")
			released, _ := result.Record().Get("l")
			movies[i].Title = title.(string)
			r, _ := strconv.ParseInt(released.(string), 10, 64)
			movies[i].Released = r
			i = i + 1
		}
		return movies, nil
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movies})
}

// GetMovieId godoc
// @Tags movie
// @Summary Get movie by id
// @Produce json
// @Param title path string true "Movie name"
// @Param released path int true "Movie released year"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie/{title}/{released} [get]
func GetMovieId(c *fiber.Ctx) error {
	n := c.Params("title")
	title, _ := url.QueryUnescape(n)
	released := string(c.Params("released"))
	movie := new(model.Movie)

	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie {title:$title,released:toInteger($released)}) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l",
			map[string]interface{}{"title": title, "released": released})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			movie.Title = result.Record().GetByIndex(0).(string)
			movie.Tagline = result.Record().GetByIndex(1).(string)
			movie.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)

			return movie, nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movie})
}

// PostMovie godoc
// @Tags movie
// @Summary Create a movies
// @Produce  json
// @Param movie body model.Movie true "Movie model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie/ [post]
func PostMovie(c *fiber.Ctx) error {
	movie := new(model.Movie)
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (m:Movie {title:$mtitle, tagline:$mtagline, released:toInteger($mreleased)}) RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l",
			map[string]interface{}{"mtitle": movie.Title, "mtagline": movie.Tagline, "mreleased": movie.Released})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			movie.Title = result.Record().GetByIndex(0).(string)
			movie.Tagline = result.Record().GetByIndex(1).(string)
			movie.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			return movie, nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movie})
}

// PutMovie godoc
// @Tags movie
// @Summary Edit a movies
// @Produce  json
// @Param title path string true "Movie name"
// @Param released path int true "Movie released year"
// @Param movie body model.Movie true "Movie model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie/{title}/{released} [put]
func PutMovie(c *fiber.Ctx) error {
	t := c.Params("title")
	title, _ := url.QueryUnescape(t)
	released := string(c.Params("released"))
	movie := new(model.Movie)
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (m:Movie {title:$title,released:toInteger($released)}) SET m.title = $mtitle, m.tagline = $mtagline, m.released = $mreleased RETURN apoc.convert.toJson(m.title) as t, apoc.convert.toJson(m.tagline) as tg, apoc.convert.toJson(m.released) as l",
			map[string]interface{}{"title": title, "released": released, "mtitle": movie.Title, "mtagline": movie.Tagline, "mreleased": movie.Released})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			movie.Title = result.Record().GetByIndex(0).(string)
			movie.Tagline = result.Record().GetByIndex(1).(string)
			movie.Released, _ = strconv.ParseInt(result.Record().GetByIndex(2).(string), 10, 64)
			return movie, nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movie})
}

// DeleteMovie godoc
// @Tags movie
// @Summary Delete a movie
// @Produce  json
// @Param title path string true "Movie name"
// @Param released path int true "Movie released year"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 401 "Unauthorized"
// @Router /movie/{title}/{released} [delete]
func DeleteMovie(c *fiber.Ctx) error {
	t := c.Params("title")
	title, _ := url.QueryUnescape(t)
	released := string(c.Params("released"))

	_, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"MATCH (m:Movie {title:$title,released:toInteger($released)}) DETACH DELETE m",
			map[string]interface{}{"title": title, "released": released})
		if err != nil {
			return nil, err
		}
		return "", nil

	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "movie deleted"})
}
