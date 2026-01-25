package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Define the shape of the data you expect to RECEIVE
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 2. Create an empty variable to hold the incoming data
		var login LoginRequest

		// 3. Decode JSON from request body into the struct
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// 4. Check the password
		if login.Password == "secret123" {
			fmt.Fprintln(w, "Login Successful")
		} else {
			fmt.Fprintln(w, "Login Failed")
		}
	})

	http.ListenAndServe(":8080", nil)
}
