package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	// Primeiro parâmentro é a rota
	// Segundo parâmetro é a função que será o handler da rota
	// Usando uma inline function
	server.GET("/ping", func(ctx *gin.Context) {
		// O que a rota irá retornar, um map
		ctx.JSON(200, gin.H{
			"message": "pong",
			"System":  "OpenSUSE",
			"Autor":   "roger437unix",
		})
	})

	server.Run(":8000")
}
