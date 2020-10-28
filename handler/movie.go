package handler

import (
	url "net/url"
	"strconv"

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
	var movie model.Movie
	m, err := repository.FindR(movie, nil)
	if err != "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

// func GetMovie(c *fiber.Ctx) error {
// 	// movies := make([]model.Movie, 10)
// 	var movie []model.Movie
// 	movie, err := repository.Find()
// 	if err != "" {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": err})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movie})
// }

// GetMovieId godoc
// @Tags movie
// @Summary Get movie by id
// @Produce json
// @Param title path string true "Movie name"
// @Param released path int64 true "Movie released year"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie/{title}/{released} [get]
func GetMovieId(c *fiber.Ctx) error {
	n := c.Params("title")
	title, _ := url.QueryUnescape(n)
	released, _ := strconv.ParseInt(c.Params("released"), 10, 64)

	var movie model.Movie
	m, err := repository.FindR(movie, map[string]interface{}{"title": title, "released": released})
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

// func GetMovieId(c *fiber.Ctx) error {
// 	n := c.Params("title")
// 	title, _ := url.QueryUnescape(n)
// 	// released := string(c.Params("released"))
// 	var movie []model.Movie
// 	movie, err := repository.Find(title, "The Matrix")
// 	if err != "" {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": movie})
// }

// PostMovie godoc
// @Tags movie
// @Summary Create a movies
// @Produce  json
// @Param movie body model.Movie true "Movie model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /movie/ [post]
func PostMovie(c *fiber.Ctx) error {
	var movie model.Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	m, err := repository.CreateR(movie)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err, "movie": m})
	}
	return c.Status(fiber.StatusOK).JSON(m)
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
	released, _ := strconv.ParseInt(c.Params("released"), 10, 64)
	var movie model.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	m, err := repository.SaveR(movie, map[string]interface{}{"title": title, "released": released})
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
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
	released, _ := strconv.ParseInt(c.Params("released"), 10, 64)
	var movie model.Movie
	m, err := repository.DeleteR(movie, map[string]interface{}{"title": title, "released": released})
	if err != "não encontrado" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": m})
}
