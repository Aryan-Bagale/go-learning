package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// 1. Define the shape of your data
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, _ := sql.Open("sqlite", "learning.db")
	defer db.Close()

	// (Setup: Insert a user so we have data)
	db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "StructMan", 25)

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		// 2. Create an instance of the struct
		var u User

		// 3. Scan directly into the struct fields
		err = rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			log.Fatal(err)
		}

		// Access data via u.Name, u.Age, etc.
		fmt.Printf("Struct User: %+v\n", u)
	}
}