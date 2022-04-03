package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// user routes go here
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetUser)
	app.Put("/user/:userId", controllers.EditUser)
	app.Delete("/user/:userId", controllers.DeleteUser)
}
