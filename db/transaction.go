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

	// 1. Start Transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Step 1: Add User A (Note: we use 'tx.Exec', not 'db.Exec')
	_, err = tx.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "User_A", 20)
	if err != nil {
		tx.Rollback() // Undo if error
		log.Fatal(err)
	}

	// 3. Step 2: Add User B
	_, err = tx.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "User_B", 21)
	if err != nil {
		tx.Rollback() // Undo if error
		log.Fatal(err)
	}

	// 4. Commit (Save all changes permanently)
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction success! Both users added.")
}