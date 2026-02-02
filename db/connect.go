package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // conncets golang to db 
)

func main(){

	db, err := sql.Open("sqlite", "learning.db")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}else {
		fmt.Println("Connected to learning.db")
	}
}