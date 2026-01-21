package main

import ("bufio"
		"os"
)

func main(){
	file, _ := os.Create("buffered.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("I am sitting in memory until flushed.")

	writer.Flush()
}