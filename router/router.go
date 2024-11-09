package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicia o router com as configurações padrão do Gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Rodando a API
	router.Run(":8080")
}