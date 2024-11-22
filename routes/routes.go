package routes

import (
	"Crud_fiber_Go/controllers"
	"Crud_fiber_Go/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Routes
	api.Post("/login", controllers.LoginUser)
	api.Post("/signup", controllers.CreateUser)
	api.Get("/users", utils.Protected, controllers.GetUsers)
	api.Delete("/users/:id", utils.Protected, controllers.DeleteUser)
	api.Put("/users/:id", utils.Protected, controllers.UpdateUser)

}
