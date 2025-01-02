/* Adicionado Migrations do models/user*/
/* Cria tabela de forma automática no banco de dados*/

package main

import (
	"fiber-project/database"
	"fiber-project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}
