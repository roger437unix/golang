package routes

import (
	"fiber-project/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Get("/api/register", controllers.Register)
	app.Post("/api/register", controllers.Register)
}
