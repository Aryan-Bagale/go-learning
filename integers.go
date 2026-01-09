package main

import "fmt"

func main() {
    var apples int = 10
    var eaten int = 4
	var price int = 2

	fmt.Println("Apples remaining:", apples - eaten)
    fmt.Println("Total cost of remaining apples:", (apples - eaten)*price)
}