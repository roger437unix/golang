package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", homepage)

	app.Listen(":3000")
}

func homepage(c *fiber.Ctx) error {
	return c.SendString("Hello, World! 😃")
}
