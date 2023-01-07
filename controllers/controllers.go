package controllers

import (
	"gin-gorm-postg/initializers"
	"gin-gorm-postg/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	// get data off req body
	var body struct {
		Title 	string
		Body 	string
		Authors string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body, Authors: body.Authors}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Index(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func Show(c *gin.Context) {
	// get the id
	id := c.Param("id") 

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// respond it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Update (c *gin.Context) {
	// Get the id off the url
	id := c.Param("id") 
	// Get the data off req body
	var body struct {
		Title 	string
		Body 	string
		Authors string
	}

	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
		Authors: body.Authors,
	})
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Delete(c *gin.Context) {

	// get id of the url
	id := c.Param("id")
	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.JSON(200, gin.H{
		"post": "Success Delete the Post",
	})
}