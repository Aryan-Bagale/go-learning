package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Build with Gin"})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		// Extract the value passed in the URL
		name := c.Param("name") 
		
		c.JSON(http.StatusOK, gin.H{
			"user": name,
		})
	})

	// NEW CODE: Query Strings
	r.GET("/search", func(c *gin.Context) {
		// Get value for 'q' from /search?q=something
		query := c.Query("q") 
		
		// Get 'page', but default to "1" if not provided
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"term": query,
			"page": page,
		})
	})

	r.Run()
}