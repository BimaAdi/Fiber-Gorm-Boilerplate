package main

import (
	"fmt"
	"os"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Intitiate from .env file (development only)
	if os.Getenv("IS_PROD") != "TRUE" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	// connect to database
	models.Initiate()

	// Fiber instance
	app := fiber.New()

	app = routes.Initiate(app)

	// start server
	fmt.Println("Start the server")
	app.Listen(":3000")
}
