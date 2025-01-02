package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	var dns = "tux:ABC123xyz@/db_loja?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Falha na conexão ao banco!!!")
	}
	fmt.Println("=> Conexão Ok!")
}
