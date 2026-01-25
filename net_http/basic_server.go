package main

import "net/http"

func main() {
    // This starts the server. 
    // It pauses the program here and listens on port 8080.
    http.ListenAndServe(":8080", nil)
	//It would be 404 because we started the server, but we didn't teach it how to reply to anything yet.
}