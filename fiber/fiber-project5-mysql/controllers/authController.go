package controllers

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.SendString("Hello, World! 😃")
}
