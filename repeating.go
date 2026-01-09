package main

import "fmt"

func main(){
	for i := 10; i > 0 ; i--{
		fmt.Println("Count:", i)
		if i == 1{
			fmt.Println("Blastoff!")
		}
	}
}