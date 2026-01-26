package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Create a default router
	r := gin.Default()

	// 2. Define a route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Build with Gin",
			"status":  "success",
		})
	})

	// 3. Start the server (defaults to port 8080)
	r.Run() 
}