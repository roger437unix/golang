package main

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := conexao()
	fmt.Println("\n=> Retorno do banco de dados MySQL:")
	fmt.Println(db, "\nType: ", reflect.TypeOf(db))

	app := fiber.New()
	app.Get("/", homepage)
	app.Listen(":3000")
}

func homepage(c *fiber.Ctx) error {
	return c.SendString("Hello, World! 😃")
}

func conexao() *gorm.DB {
	var dns = "tux:ABC123xyz@/db_loja?charset=utf8mb4&parseTime=True&loc=Local"
	var v = "\n*** Sem conexão ao banco de dados!!!"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(v)
	}
	return db
}
