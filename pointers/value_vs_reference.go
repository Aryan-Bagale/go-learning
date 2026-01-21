package main

import "fmt"

// 1. Primitive: NEEDS '*' to modify
func changeNumber(n *int) {
    *n = 500
}

// 2. Slice: NO '*' needed to modify content
// (Has a hidden pointer to the array)
func changeSlice(s []int) {
    s[0] = 500
}

// 3. Map: NO '*' needed to modify content
// (Is a reference/hidden pointer itself)
func changeMap(m map[string]int) {
    m["score"] = 500
}

func deleteKey(m map[string]int){
	delete(m, "score")
}

func main() {
    // --- Primitive ---
    num := 10
    changeNumber(&num) // Must use &
    fmt.Println("Number:", num)

    // --- Slice ---
    list := []int{10, 20}
    changeSlice(list)  // Just pass variable
    fmt.Println("Slice: ", list)

    // --- Map ---
    data := map[string]int{"score": 10}
    changeMap(data)    // Just pass variable
    fmt.Println("Map:   ", data)

	deleteKey(data)
	fmt.Println("After delete key map:", data)
}