package handler

import (
	"net/url"

	"github.com/guimaraaes/golang_fiber_with_neo4j/repository_algorithms"

	"github.com/gofiber/fiber/v2"
)

// Algorithms godoc
// @Tags algorithms
// @Summary Show centrality nodes
// @Produce  json
// @Param node path string true "Node type"
// @Param relationship path string true "Relationship projection"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /algo_centrality/{node}/{relationship} [get]
func Centrality(c *fiber.Ctx) error {
	n := c.Params("node")
	node, _ := url.QueryUnescape(n)
	r := c.Params("relationship")
	relationship, _ := url.QueryUnescape(r)

	res, err := repository_algorithms.Centrality(node, relationship)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": res})
}

// Algorithms godoc
// @Tags algorithms
// @Summary Show comunity nodes
// @Produce  json
// @Param node path string true "Node type"
// @Param relationship path string true "Relationship projection"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /algo_community/{node}/{relationship} [get]
func Community(c *fiber.Ctx) error {
	n := c.Params("node")
	node, _ := url.QueryUnescape(n)
	r := c.Params("relationship")
	relationship, _ := url.QueryUnescape(r)

	res, err := repository_algorithms.Community(node, relationship)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": res})
}

// Algorithms godoc
// @Tags algorithms
// @Summary Show path nodes
// @Produce  json
// @Param node path string true "Node type"
// @Param relationship path string true "Relationship projection"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /algo_path/{node}/{relationship} [get]
func Path(c *fiber.Ctx) error {
	n := c.Params("node")
	node, _ := url.QueryUnescape(n)
	r := c.Params("relationship")
	relationship, _ := url.QueryUnescape(r)

	res, err := repository_algorithms.Path(node, relationship)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": res})
}

// Algorithms godoc
// @Tags algorithms
// @Summary Show page rank nodes
// @Produce  json
// @Param node path string true "Node type"
// @Param relationship path string true "Relationship projection"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /algo_pagerank/{node}/{relationship} [get]
func PageRank(c *fiber.Ctx) error {
	n := c.Params("node")
	node, _ := url.QueryUnescape(n)
	r := c.Params("relationship")
	relationship, _ := url.QueryUnescape(r)

	res, err := repository_algorithms.PageRank(node, relationship)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": res})
}
