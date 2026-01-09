package main

import "fmt"

func main(){

	menu:= map[string]int{
		"Burger": 10,
		"Fries": 5,
	}
	menu["Soda"] = 3

	fmt.Println("Whole Menu", menu)
	fmt.Println("Price of Fries:", menu["Fries"] )
	fmt.Println("Price of Soda:", menu["Soda"])

}