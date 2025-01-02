/* Alterando a função conexão para retornar dois argumentos */

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	_, err := conexao()
	if err != nil {
		fmt.Println("Falha na conexão ao banco!!!")
	}
	// fmt.Println(db, "\nType: ", reflect.TypeOf(db))

	app := fiber.New()
	app.Get("/", homepage)
	app.Listen(":3000")
}

func homepage(c *fiber.Ctx) error {
	return c.SendString("Hello, World! 😃")
}

func conexao() (*gorm.DB, error) {
	var dns = "tux:ABC123xyz@/db_loja?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	return db, err
}
