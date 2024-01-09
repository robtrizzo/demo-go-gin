package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetSearchPage(c *gin.Context) {
	render(c, gin.H{}, "search.html")
}

func Search(c *gin.Context) {
	render(c, gin.H{}, "search_results.html")
}
