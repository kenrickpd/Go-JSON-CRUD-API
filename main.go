package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenrickpd/Go-JSON-CRUD-API/controllers"
	"github.com/kenrickpd/Go-JSON-CRUD-API/initializers"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
