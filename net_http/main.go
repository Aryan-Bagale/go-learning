package main

import (
	"fmt"
	"net/http"
)

func main() {
	// The function "listens" to the request (r)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		// 1. LISTEN: Check what they typed in the URL (The Input)
		whatUserTyped := r.URL.Path

		// 2. THINK: Decide what to say
		if whatUserTyped == "/pizza" {
			// 3. SPEAK: Talk into the mouthpiece (The Output)
			fmt.Fprint(w, "Here is your Pepperoni Pizza!")
		} else {
			fmt.Fprint(w, "We only serve /pizza here.")
		}
	})

	http.ListenAndServe(":8080", nil)
}