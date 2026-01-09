package main

import "fmt"


func main(){
	var battery int = 100
	if battery < 20 {
		fmt.Println("Warning: Low Battery!")
	} else if battery == 100 {
		fmt.Println("Fully Charged!")
	}else {
		fmt.Println("System Normal.")
	}
}