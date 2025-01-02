package controllers

import (
	"fiber-project/models"

	"github.com/gofiber/fiber/v2"
)

// retorna o que será exibido no navegador
func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "tux",
	}
	user.LastName = "unix"
	user.Email = "tux@kernel.org"
	return c.JSON(user)
}
