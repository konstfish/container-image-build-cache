package main

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func main() {
	Router = gin.Default()

	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong2",
		})
	})

	Router.Run(":8080")
}
