package main

import (
	"github.com/gin-gonic/gin"

	"github.com/robtrizzo/semaphore-demo-go-gin/handlers"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	// Handle Index
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
	// Serve the application
	router.Run()
}
