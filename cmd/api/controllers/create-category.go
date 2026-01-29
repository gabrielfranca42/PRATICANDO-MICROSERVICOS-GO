package controllers

import {
	"github.com/gin-gon/gin"
}


type CreateCategoryInput struct{
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context){
	var body CreateCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJson(htpp.StatusBadRequest,
		    gin.H{
			  "sucess": false,
			  "error": err,Error(),
		})
	return 

	}

	useCase := use_cases.NewCreateCategoryUseCase{}

	useCase.Execute()
}
