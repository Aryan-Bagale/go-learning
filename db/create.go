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

	sts := `
	CREATE TABLE IF NOT EXISTS users(
			id INTEGER 	PRIMARY KEY,
			name TEXT,
			age INTEGER
	);`

	_, err := db.Exec(sts)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Table 'users' created successfully.")
}