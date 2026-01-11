package main

import "fmt"

func main(){
	score := 100
	name := "Batman"

	p := &score // storing the address of variable in another variable, this is basically called pointer
	b := &name

	fmt.Println("Score:", score)
	fmt.Println("Address of score is:", p)

	fmt.Println("Name:", name)
	fmt.Println("Address of score is:", p)

	*p = 200 // here we  took the stored address of score which was in p and just give it new value 
	fmt.Println("New Score is:", score)

	*b = "Robin"
	fmt.Println("New Name:", name)

}