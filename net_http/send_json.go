package main

import (
	"encoding/json"
	"net/http"
)

type UserData struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdmin bool `json:"is_admin"`
	Email string `json:"email_address"`
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		//Create an instance of the struct
		user := UserData{
			Name: "Aryan",
			Age: 20,
			IsAdmin: true,
			Email: "example@gmail.com",
		}

		// Set the "Header" so the browser knows JSON is coming
		w.Header().Set("Content-Type", "application/json")


		// Encode the struct directly
		json.NewEncoder(w).Encode(user)
		//It automatically converts 25 to a number and true to a boolean in the JSON output.
	})

	http.ListenAndServe(":8080", nil)
}