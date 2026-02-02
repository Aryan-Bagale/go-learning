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

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next(){
		var id int
		var name string
		var age int

		err = rows.Scan(&id, &name, &age)
		if err != nil{
			log.Fatal(err)
		}

		fmt.Printf("User %s is %d years old\n",name, age)
	}
}