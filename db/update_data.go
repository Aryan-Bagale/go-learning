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

	// 1. Prepare the Update statement
	updateSQL := "UPDATE users SET age = ? WHERE name = ?"

	// 2. Execute it (Set Alice's age to 99)
	result, err := db.Exec(updateSQL, 99, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Check how many rows were updated
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Updated %d row(s).\n", rowsAffected)
}