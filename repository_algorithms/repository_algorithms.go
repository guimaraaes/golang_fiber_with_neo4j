package repository_algorithms

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
	_ "github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
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
	r := c.Params("relationshio")
	relationship, _ := url.QueryUnescape(r)

	fmt.Println(node)
	fmt.Println(relationship)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "s"})
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

	fmt.Println(node)
	fmt.Println(relationship)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "s"})
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

	fmt.Println(node)
	fmt.Println(relationship)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "s"})
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

	// fmt.Println(node)
	// fmt.Println(relationship)

	query := "CALL gds.pageRank.stream({ nodeProjection: [$node], " +
		"relationshipProjection: [$rel], " +
		"maxIterations: 1, dampingFactor: 0.85}) " +
		"YIELD nodeId, score " +
		"RETURN apoc.convert.toJson(gds.util.asNode(nodeId).name) ORDER BY score DESC"
	m := map[string]interface{}{"node": node, "rel": relationship}
	res, _ := utils.QueryCall(query, m)

	// fmt.Println(query)
	// fmt.Println(res)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": res})
}
