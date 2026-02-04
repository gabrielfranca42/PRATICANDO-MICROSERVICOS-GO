package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	controllers "github.com/gabrielfranca42/go-microservices/cmd/api/controllers"
	repo "github.com/gabrielfranca42/go-microservices/internal/entities_models/repository"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("ENVIRONMENT")

	if env == "local" {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// health
	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	// cria repository concreto
	categoryRepository := repo.NewInMemoryCategoryRepository()

	// registra rota direto (sem CategoryRoutes)
	r.POST("/categories", func(c *gin.Context) {
		controllers.CreateCategory(c, categoryRepository)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
