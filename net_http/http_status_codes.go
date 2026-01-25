package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
        
        // 1. Handle Bad JSON
		if err != nil {
            // w.WriteHeader sets the status code
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprint(w, "Invalid JSON data")
			return // STOP here! Don't continue.
		}

		if u.Password == "secret123" {
            // 2. Handle Success
			w.WriteHeader(http.StatusOK) // 200
			fmt.Fprint(w, "Login Successful!")
		} else {
            // 3. Handle Wrong Password
			w.WriteHeader(http.StatusUnauthorized) // 401
			fmt.Fprint(w, "Login Failed: Wrong Password")
		}
	})

	http.ListenAndServe(":8080", nil)
}