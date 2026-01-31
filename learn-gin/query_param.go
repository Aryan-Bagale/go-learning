package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	// Create a Gin router with logger and recovery middleware
	r := gin.Default()

	// ----------------------------
	// ROOT ROUTE
	// ----------------------------
	// Handles: GET /
	r.GET("/", func(c *gin.Context) {

		// Send a JSON response with HTTP 200
		c.JSON(http.StatusOK, gin.H{
			"message": "Build with Gin",
		})
	})

	// ----------------------------
	// PATH PARAMETER ROUTE
	// ----------------------------
	// Handles: GET /user/:name
	// Example: /user/alice
	r.GET("/user/:name", func(c *gin.Context) {

		// c.Param("name") extracts the value from the URL path
		// /user/alice -> name = "alice"
		name := c.Param("name")

		// Return the extracted value as JSON
		c.JSON(http.StatusOK, gin.H{
			"user": name,
		})
	})

	// ----------------------------
	// QUERY STRING ROUTE: SEARCH
	// ----------------------------
	// Handles: GET /search?term=go&page=2

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

		// c.Query("term") reads a query parameter
		// /search?term=go -> term = "go"
		// If missing, it returns ""
		term := c.Query("term")

		// c.DefaultQuery("page", "1")
		// If page is missing, "1" is used automatically
		page := c.DefaultQuery("page", "1")

		// Send query values back as JSON
		c.JSON(http.StatusOK, gin.H{
			"search_term": term,
			"page_number": page,
		})
	})

	// ----------------------------
	// QUERY STRING ROUTE: FILTER
	// ----------------------------
	// Handles: GET /filter?type=books&sort=price
	r.GET("/filter", func(c *gin.Context) {

		// Reads "type" from the query string
		// /filter?type=books -> "books"
		filterType := c.Query("type")

		// Reads "sort", defaults to "price" if missing
		sortBy := c.DefaultQuery("sort", "price")

		// Return filter values as JSON
		c.JSON(http.StatusOK, gin.H{
			"type": filterType,
			"sort": sortBy,
		})
	})

	// Start the HTTP server on :8080
	r.Run()
}
