package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "learning.db")
	defer db.Close()

	// Setup: Insert Bob (since we deleted everyone else)
	db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Bob", 40)

	var name string
	var age int

	// 1. QueryRow executes and prepares to Scan in one chain
	err := db.QueryRow("SELECT name, age FROM users WHERE name = ?", "Bob").Scan(&name, &age)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found single user: %s, Age: %d\n", name, age)
}