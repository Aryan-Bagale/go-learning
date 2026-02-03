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

	// 1. Prepare the Delete statement
	deleteSQL := "DELETE FROM users WHERE name = ?"

	// 2. Execute it (Delete 'Alice')
	result, err := db.Exec(deleteSQL, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Verify deletion
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Deleted %d row(s).\n", rowsAffected)
}