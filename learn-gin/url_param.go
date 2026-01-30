package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Build with Gin"})
	})

    // --- STEP 2: DYNAMIC PARAMETERS ---

	// Without Gin:
	// http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
	//     // You have to manually parse the URL string to find the name
	//     name := strings.TrimPrefix(r.URL.Path, "/user/") 
	// })

	// With Gin:
	// Use ':' to denote a dynamic parameter
	r.GET("/user/:name", func(c *gin.Context) {
		
		// Gin Helper: c.Param()
		// Automatically extracts the value at the ':name' position
		name := c.Param("name") 
		
		c.JSON(http.StatusOK, gin.H{
			"result": name,
		})
	})

	// Dynamic parameter: number (validated)
	r.GET("/id/:number", func(c *gin.Context) {
		numberStr := c.Param("number")

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id must be a number",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"your_id": number,
		})
	})

	// Start server on :8080
	r.Run(":8080")
}