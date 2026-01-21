package main

import "fmt"

type Player struct{
	Name string
	Level int
	Alive bool
	Class string
}

func main(){
	p1 := Player{
		Name: "Hero",
		Level: 10,
		Alive: true,
		Class: "Rogue",
	}

	fmt.Println("Player Name:", p1.Name)
	fmt.Println("Full Stats:", p1)
	fmt.Println("Class:", p1.Class)

}