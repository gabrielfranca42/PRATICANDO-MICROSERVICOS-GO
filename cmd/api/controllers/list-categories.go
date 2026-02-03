package controllers

import (
	"net/http"

	"github.com/gabrielfranca42/go-microservices/internal/entities_models/repository"
	"github.com/gabrielfranca42/go-microservices/services"
	"github.com/gin-gonic/gin"
)

func ListCategory(ctx *gin.Context, repository repository.ICategoryRepository) {

	useCase := services.NewListCategoryUseCase(repository)

	category, err := useCase.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true, "category": category,
	})
}
