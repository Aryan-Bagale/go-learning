package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We are sending HTML tags inside the string
		fmt.Fprint(w, "<h1>Welcome to my Go Site!</h1>")
		fmt.Fprint(w, "<p>This is a paragraph of text.</p>")
	})

	http.ListenAndServe(":8080", nil)
}