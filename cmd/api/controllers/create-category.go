package controllers

import (
	"net/http"

	"github.com/gabrielfranca42/go-microservices/internal/entities_models/repository"
	"github.com/gabrielfranca42/go-microservices/services"
	"github.com/gin-gonic/gin"
)

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repository.ICategoryRepository) {
	var body CreateCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	useCase := services.NewCreateCategoryUseCase(repository)

	err := useCase.Execute(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
