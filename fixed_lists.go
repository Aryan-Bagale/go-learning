package main

import "fmt"

func main(){
	var inventory [3]string
	inventory[0] = "Sword"
	inventory[1] = "Shield"
	inventory[2] = "Potion"

	inventory[1] = "Broken Shield"

	fmt.Println("Item 3:", inventory[2])
	fmt.Println("Full Inventory:", inventory)
}