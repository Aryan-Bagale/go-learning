package main

import (
	"fmt"
	"net/http"
)

func main() {
	// This captures ONLY exact matches for "/bye"
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Goodbye!")
	})

	// This captures "/" AND everything else (like /pizza, /random, /test)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "You hit the catch-all! Path: ", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}