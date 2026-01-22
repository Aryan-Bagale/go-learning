package main

import "fmt"

func Describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i,i)
}

func main(){

	Describe("Hello, World")

	Describe(42)

	Describe(struct{Name string}{"Anonymous"})

	data := map[string]interface{}{
		"age": 25,
		"name": "Alice",
		"active": true,
	}

	Describe(data)

}