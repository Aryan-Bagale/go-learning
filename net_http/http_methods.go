package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if r.Method == "GET"{
			fmt.Fprint(w, "You requested to READ data (GET)")
		} else if r.Method == "POST" {
			fmt.Fprint(w, "You requested to WRITE data (POST).")

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}