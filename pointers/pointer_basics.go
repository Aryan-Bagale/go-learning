package main

import "fmt"

func main(){

	name := "Go"

	fmt.Println("Name", name)

	var p *string = &name
	// same as : p := &name

	fmt.Println("Address in memory:", p)
	fmt.Println("Value via pointer:", *p)
	
	*p = "Golang"

	fmt.Println("New Name:",name)
}
