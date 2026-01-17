package main

import "os"

func main() {
    // 1. Open with specific flags: Append, Create if missing, Write Only
    file, err := os.OpenFile("greetings.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // 2. Add new line
    file.WriteString("\nI am adding this line.")
}