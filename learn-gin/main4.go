package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// --- STEP 4: DEFINE DATA STRUCTURE ---
// We need a Struct to define what the JSON looks like.
// `json:"title"` tells Go to map the JSON key "title" to this field.
type NewGame struct {
	Title string `json:"title"`
	Genre string `json:"genre"`
}

func main() {
	r := gin.Default()

	// --- STEP 1: BASIC ROUTE ---
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Build with Gin"})
	})

	// --- STEP 2: DYNAMIC PARAMETERS ---
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"user": name})
	})

	// --- STEP 3: QUERY STRINGS ---
	r.GET("/search", func(c *gin.Context) {
		term := c.Query("term")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{"term": term, "page": page})
	})

	// --- STEP 4: POST REQUESTS (NEW) ---
	
	// Without Gin:
	// http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
	//     var game NewGame
	//     // 1. Create decoder, 2. Decode, 3. Handle Error
	//     if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
	//         http.Error(w, err.Error(), http.StatusBadRequest)
	//         return
	//     }
	//     ...
	// })

	// With Gin:
	r.POST("/games", func(c *gin.Context) {
		var newGame NewGame

		// Gin Helper: c.ShouldBindJSON()
		// Automatically reads the request body and maps it to the struct.
		// It returns an error if the JSON format is wrong.
		if err := c.ShouldBindJSON(&newGame); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return the data to confirm we received it
		c.JSON(http.StatusCreated, gin.H{
			"message": "Game added",
			"data":    newGame,
		})
	})

	r.Run()
}