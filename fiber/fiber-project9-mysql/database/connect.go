package database

import (
	"fiber-project/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dns = "tux:ABC123xyz@/db_loja?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Falha na conexão ao banco!!!")
	}

	DB = conn

	// Irá criar tabela de forma automática baseado na struct "models.User"
	conn.AutoMigrate(&models.User{})

	fmt.Println("=> Conexão Ok!")
}
