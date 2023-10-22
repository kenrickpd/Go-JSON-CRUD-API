package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kenrickpd/Go-JSON-CRUD-API/initializers"
	"github.com/kenrickpd/Go-JSON-CRUD-API/models"
)

func UserCreate(c *gin.Context) {
	var attribute struct {
		Name  string
		Age   int
		Email string
	}

	c.Bind(&attribute)

	//Create new posts
	user := models.User{Name: attribute.Name, Age: attribute.Age, Email: attribute.Email}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return nya
	c.JSON(200, gin.H{
		"user": user,
	})
}
