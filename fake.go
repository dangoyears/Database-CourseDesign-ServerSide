package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Human struct {
	Username string `form:"name"`
	Password string `form:"pass"`
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/login", func(c *gin.Context) {
		var human Human

		if (c.ShouldBind(&human)) == nil && human.Password == "dangoyears" {
			c.JSON(200, gin.H{
				"token": "db4de92f933092b1f1f6665a742b7192a5eced77",
			})
		} else {
			c.JSON(200, gin.H{
				"token": "",
			})
		}

	})

	r.Run("localhost:12323")
}
