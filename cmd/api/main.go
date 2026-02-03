package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	c_routes "github.com/gabrielfranca42/go-microservices/cmd/api/routes"
	infraDb "github.com/gabrielfranca42/go-microservices/internal/infra/db"
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

	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	db, err := infraDb.InitDb()

	if err != nil {
		panic(err)
	}

	infraDb.MigrateModels(db)

	c_routes.CategoryRoutes(r, db)

	r.Run(":8080")
}
