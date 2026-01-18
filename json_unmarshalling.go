
//unmarshalling means converting JSON into a Go data structure.

package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int
	Rating float64
}

func main() {
	// 1. Simulating raw JSON data (use backticks ` for strings)
	jsonString := `{"Title": "The Matrix", "Year": 1999, "Rating":8.7}`

	// 2. Create an empty variable to hold the result
	var m Movie

	// 3. Convert JSON to Struct
	// distinct note: we must pass the POINTER (&m) so Go can update it
	json.Unmarshal([]byte(jsonString), &m)

	// 4. Print the result
	fmt.Println("Movie Info:", m.Title, m.Year, m.Rating)
}