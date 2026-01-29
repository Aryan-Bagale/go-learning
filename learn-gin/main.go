package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// --- STEP 1: ROUTER SETUP ---
	
	// Without Gin: 
	// mux := http.NewServeMux() (You have to create a multiplexer manually)
	
	// With Gin:
	r := gin.Default() // Creates router + Logger + Crash Recovery


	// --- STEP 2: DEFINING A ROUTE ---

	// Without Gin: 
	// mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	
	// With Gin:
	r.GET("/ping", func(c *gin.Context) {

		// --- STEP 3: SENDING JSON ---

		// Without Gin:
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(map[string]string{"message": "pong"})

		// With Gin:
		// (Handles Headers, Status Code, and Encoding in one line)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/info", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status":"active",
		})
	})

	// --- STEP 4: STARTING SERVER ---

	// Without Gin:
	// http.ListenAndServe(":8080", mux)

	// With Gin:
	r.Run() // Defaults to :8080
}