package handler

import (
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
	greeting, err := database.Neo4jDS.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (a:Greeting) SET a.message = 'a' RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": greeting.(string)})

}

// GetMovieId godoc
// @Tags movie
// @Summary Get movie by id
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router{id} /movie [get]
func GetMovieId(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
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
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

// PutMovie godoc
// @Tags movie
// @Summary Edit a movies
// @Produce  json
// @Param id path int true "Movie ID"
// @Param movie body model.Movie true "Movie model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie/{id} [put]
func PutMovie(c *fiber.Ctx) error {
	var movie model.Movie

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": movie})
}

// DeleteMovie godoc
// @Tags movie
// @Summary Delete a movies
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 401 "Unauthorized"
// @Router /movie/{id} [delete]
func DeleteMovie(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}
