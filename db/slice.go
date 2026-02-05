package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, _ := sql.Open("sqlite", "learning.db")
	defer db.Close()

	// Setup: Add a few more people
	db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Dave", 34)
	db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Eve", 22)

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 1. Create an empty slice to hold the users
	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatal(err)
		}
		// 2. Append the user to the slice
		users = append(users, u)
	}

	// 3. Loop through the slice (in memory)
	fmt.Println("--- User List ---")
	for _, u := range users {
		fmt.Printf("%s is %d\n", u.Name, u.Age)
	}
}