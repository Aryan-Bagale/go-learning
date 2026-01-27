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

    // NEW CODE: Dynamic Route using ':'
	r.GET("/user/:name", func(c *gin.Context) {
		// Extract the value passed in the URL
		name := c.Param("name") 
		
		c.JSON(http.StatusOK, gin.H{
			"user": name,
		})
	})

	r.Run()
}