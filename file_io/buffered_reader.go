package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // 1. Open the file
    file, err := os.Open("greetings.txt") // We use Open, not Create
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // 2. Create the Scanner
    scanner := bufio.NewScanner(file)

    // 3. Loop through every line
    // scanner.Scan() returns true as long as there is a new line
    for scanner.Scan() {
        fmt.Println("Read line:", scanner.Text())
    }
}