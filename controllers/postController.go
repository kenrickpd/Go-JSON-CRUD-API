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

	//return nya
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//send respond
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsFind(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//find a posts by id
	var post models.Post
	initializers.DB.First(&post, id)

	//send respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id by url
	id := c.Param("id")

	// get data from the body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post
	var post models.Post
	initializers.DB.First(&post, id)

	//update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//send respond
	c.JSON(200, gin.H{
		"post": post,
	})
}
