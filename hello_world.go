package main

import "fmt"

func main(){
	var name string
	fmt.Print("Type your name: ")
	fmt.Scan(&name)
	fmt.Printf("System Online: Hello, %s!\n", name)
	fmt.Println("Ready to code.")
}