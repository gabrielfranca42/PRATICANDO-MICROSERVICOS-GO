package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/healty", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"rodou esa porra ": true,
		})
	})
	r.Run(":8080")
}
