
// Marshalling means converting a Go data structure into JSON,
package main

import(
	"encoding/json"
	"fmt"
)

// 1. Define the structure of your data
type Movie struct{
	Title string
	Director string
	Year int
}

func main(){

	// 2. Create an instance of the struct
	m := Movie{Title: "Inception", Director:"yiu", Year: 2010}
	
	// 3. Convert struct to JSON bytes
	// We use _ to skip error handling for this basic example
	jsonData, _ := json.Marshal(m)

	// 4. Print the JSON string
	fmt.Println(string(jsonData))
}
