package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenrickpd/Go-JSON-CRUD-API/initializers"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
