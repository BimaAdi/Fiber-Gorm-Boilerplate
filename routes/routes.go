package routes

import (
	"fmt"

	userroutes "github.com/BimaAdi/fiberPostgresqlBoilerPlate/routes/userRoutes"
	"github.com/gofiber/fiber/v2"
)

// Initiate all routes
func Initiate(app *fiber.App) *fiber.App {
	fmt.Println("Initiate Routes")

	userRoutes := app.Group("/user")
	userRoutes.Get("/", userroutes.GetAllUserRoute)
	userRoutes.Get("/:id", userroutes.GetDetailUserRoute)
	userRoutes.Post("/", userroutes.CreateUserRoute)
	userRoutes.Put("/:id", userroutes.UpdateUserRoute)
	userRoutes.Delete("/:id", userroutes.DeleteUserRoute)

	return app
}
