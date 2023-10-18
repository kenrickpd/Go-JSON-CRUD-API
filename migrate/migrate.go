package main

import (
	"github.com/kenrickpd/Go-JSON-CRUD-API/initializers"
	"github.com/kenrickpd/Go-JSON-CRUD-API/models"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
