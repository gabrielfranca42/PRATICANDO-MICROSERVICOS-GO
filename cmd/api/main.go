package main

import (
	"github.com/gabrielfranca42/go-microservices/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healty", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"rodou essa porra": true,
		})
	})

	// registra rotas de categoria
	r.POST("/categories", controllers.CreateCategory)

	r.Run(":8080")
}
