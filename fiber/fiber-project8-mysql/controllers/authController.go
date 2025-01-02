package controllers

import (
	"fiber-project/models"

	"github.com/gofiber/fiber/v2"
)

// gerencia a requisição que será enviada ou recebida
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)

		return c.JSON(fiber.Map{
			"message": "Senhas diferentes",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  data["password"],
	}
	return c.JSON(user)
}
