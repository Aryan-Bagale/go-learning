package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {


		// Dig into the 'r' box to find the 'name' value
		// This looks for ?name=... in the URL
		name := r.URL.Query().Get("name")

		if name == ""{
			fmt.Fprint(w, "Hello, Stranger!")
		} else{ fmt.Fprint(w, "Hello, " + name + "!")}
	})

	http.ListenAndServe(":8080", nil)
}