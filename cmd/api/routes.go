package main

import (
	"github.com/gabrielfranca42/go-microservices/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(route *gin.Engine) {
	CategoryRoutes := route.Group("/categories")

	CategoryRoutes.POST("/", controllers.CreateCategory)
}
