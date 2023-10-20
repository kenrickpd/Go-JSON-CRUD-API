package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kenrickpd/Go-JSON-CRUD-API/initializers"
	"github.com/kenrickpd/Go-JSON-CRUD-API/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create new posts
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
