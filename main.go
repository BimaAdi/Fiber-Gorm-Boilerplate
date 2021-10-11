package main

import (
	"fmt"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// connect to database
	models.Initiate()

	// Fiber instance
	app := fiber.New()

	app = routes.Initiate(app)

	// start server
	fmt.Println("Start the server")
	app.Listen(":3000")
}
