package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/robtrizzo/semaphore-demo-go-gin/models"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}
