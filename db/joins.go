package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type PostInfo struct {
	Title    string
	UserName string
}

func main() {
	db, _ := sql.Open("sqlite", "learning.db")
	defer db.Close()

	// 1. Create table with a 'Foreign Key' concept (user_id)
	db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY,
		title TEXT,
		user_id INTEGER
	)`)

	// 2. Insert a post for User ID 1 (Ensure User ID 1 exists!)
	// If ID 1 is missing, change the '1' to an ID that exists in your DB
	db.Exec("INSERT INTO posts (title, user_id) VALUES (?, ?)", "Golang is great", 1)

	// 3. JOIN Query: Get Post Title AND the User's Name
	query := `
		SELECT posts.title, users.name 
		FROM posts 
		INNER JOIN users ON posts.user_id = users.id
		WHERE posts.user_id = ?`

	var p PostInfo
	// Get post info for User ID 1
	err := db.QueryRow(query, 1).Scan(&p.Title, &p.UserName)
	if err != nil {
		log.Println("Error (User ID 1 might not exist):", err)
	} else {
		fmt.Printf("Post: '%s' | Author: %s\n", p.Title, p.UserName)
	}
}