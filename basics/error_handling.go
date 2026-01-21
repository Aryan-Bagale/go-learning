package main

import (
	"errors"
	"fmt"
)

func Divide(a,b int) (int, error){
	if b== 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a/b, nil
}

func main(){
	result, err := Divide(10,2)

	if err != nil {
		fmt.Println("Error detected:", err)
	} else {
		fmt.Println("Success! Result:", result)
	}
	fmt.Println("Calculation complete")
}