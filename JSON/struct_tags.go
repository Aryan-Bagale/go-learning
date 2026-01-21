package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"released_at"`
	Hidden string `json:"-"`
	Director string `json:"directed_by"`
}

func main(){
	m := Movie{
		Title : "Inception",
		Year: 2010,
		Hidden: "SecretDatabaseID",
		Director: "Christopher Nolan",
	}

	jsonData, _ := json.Marshal(m)

	fmt.Println(string(jsonData))
}