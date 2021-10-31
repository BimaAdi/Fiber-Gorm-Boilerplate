package routes

import (
	"fmt"

	authroutes "github.com/BimaAdi/fiberGormBoilerPlate/routes/authRoutes"
	userroutes "github.com/BimaAdi/fiberGormBoilerPlate/routes/userRoutes"
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

	authRoutes := app.Group(("/auth"))
	authRoutes.Post("/login", authroutes.LoginRoute)
	authRoutes.Post("/logout", authroutes.LogoutRoute)

	return app
}
