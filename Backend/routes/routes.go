package routes

import (
	"Crud_fiber_Go/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Routes
	api.Post("/login", controllers.LoginUser)
	api.Post("/signup", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)
	api.Delete("/users/:id", controllers.DeleteUser)
	api.Put("/users/:id", controllers.UpdateUser)

}
