package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main(){
	db, _ := sql.Open("sqlite", "learning.db")
	defer db.Close()

	insertSQL := `INSERT INTO users (name,age) VALUES (?, ?)`

	_, err := db.Exec(insertSQL, "Aryan", 20)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted user successfully.")
}