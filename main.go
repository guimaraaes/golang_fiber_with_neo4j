package main

import (
	// "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/guimaraaes/golang_fiber_with_neo4j/database"
	_ "github.com/guimaraaes/golang_fiber_with_neo4j/docs"
	"github.com/guimaraaes/golang_fiber_with_neo4j/router"
	"github.com/qinains/fastergoding"
)

// @title fiber aplication with neo4j
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	fastergoding.Run()
	database.ConnectionNeo4j()
	app := fiber.New()
	router.Routes(app)

	app.Listen(":3000")
}
