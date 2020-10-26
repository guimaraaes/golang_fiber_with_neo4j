package handler

import (
	url "net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/guimaraaes/golang_fiber_with_neo4j/repository"
)

// GetMovie godoc
// @Tags movie
// @Summary Show all movies
// @Produce  json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie [get]
func GetMovie(c *fiber.Ctx) error {
	// movies := make([]model.Movie, 10)
	var movie []model.Movie
	movie, err := repository.Find()
	if err != "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movie})
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
	var movie []model.Movie
	movie, err := repository.Find(title, released)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
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
	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	m, err := repository.Create(movie)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err, "movie": m})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": m})
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
	m, err := repository.Save(title, released, movie)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err, "movie": movie})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": m})
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
	err := repository.Delete(title, released)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "movie excluído"})
}
