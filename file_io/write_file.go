package main

import (
	"os"
)

func main() {

	file, err := os.Create("greetings.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	//this will execute at end of main function 

	_, err = file.WriteString("Hello, World!\nWelcome to Go programming.") 
	// here the first variable we omitted beacuse it gave number of bytes actually written to the file
	if err != nil {
		panic(err)
	}
}
