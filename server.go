package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"greeting": gin.H{
				"en": "Hello from Recepta!",
				"fr": "Bonjour de Recepta!",
				"ko": "Recepta에서 안녕하세요!",
				"zh": "来自Recepta的问候!",
				"ru": "Привет от Recepta!",
				"es": "¡Hola desde Recepta!",
				"de": "Hallo von Recepta!",
				"ja": "Receptaからこんにちは!",
				"ar": "مرحبًا من Recepta!",
				"pt": "Olá do Recepta!",
				"it": "Ciao da Recepta!",
			},
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}