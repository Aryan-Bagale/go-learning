package main

import "fmt"

type Clicker interface {
	Click()
}

type Counter struct{
	Value int
}

func (c *Counter) Click(){
	c.Value++
	fmt.Println("Value is:", c.Value)
}

func main(){

	myCounter := Counter{Value : 0}

	fmt.Println("Start Value:", myCounter.Value)

	var button Clicker = &myCounter

	button.Click()
	button.Click()

	fmt.Println("Final check:", myCounter.Value)

}