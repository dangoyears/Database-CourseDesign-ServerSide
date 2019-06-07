package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"token": "db4de92f933092b1f1f6665a742b7192a5eced77",
		})
	})

	r.Run("localhost:12323")
}
