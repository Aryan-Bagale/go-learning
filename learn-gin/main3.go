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

	// --- STEP 3: QUERY STRINGS ---

	// Without Gin:
	// http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
	//     // 1. You must parse the query string map manually
	//     q := r.URL.Query()
	//     term := q.Get("term")
	//     // 2. Default values require manual "if" checks
	//     page := q.Get("page")
	//     if page == "" { page = "1" }
	// })

	// With Gin:
	r.GET("/search", func(c *gin.Context) {
		
		// Gin Helper: c.Query()
		// Gets the value for "term" from the URL
		term := c.Query("term")

		// Gin Helper: c.DefaultQuery()
		// Gets "page", but returns "1" if the user didn't send it
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"search_term": term,
			"page_number": page,
		})
	})

	r.Run()
}