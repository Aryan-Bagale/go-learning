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

	// 1. Prepare the statement ONCE
	stmt, err := db.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Close the statement when done

	// 2. Execute it multiple times with different data
	// Notice we call stmt.Exec, not db.Exec
	_, _ = stmt.Exec("Prep_User1", 50)
	_, _ = stmt.Exec("Prep_User2", 51)
	_, _ = stmt.Exec("Prep_User3", 52)

	fmt.Println("Bulk insert complete using Prepared Statement.")
}