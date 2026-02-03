package main

import (
	"github.com/gabrielfranca42/go-microservices/cmd/api/controllers"
	"github.com/gabrielfranca42/go-microservices/internal/entities_models/repository"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(route *gin.Engine) {
	CategoryRoutes := route.Group("/categories")

	inMemoryCategoryRepository := repository.NewInMemoryCategoryRepository()

	CategoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})
}
