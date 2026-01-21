package main

import "fmt"

type User struct {
	username string
	isActive bool
	level int
}

func main(){

	user1 := User{username:"dev_one", isActive: false, level: 1}

	ptr := &user1

	ptr.isActive = true

	ptr.level = 5

	fmt.Println("original struct updated:", user1)
}