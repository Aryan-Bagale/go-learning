package main

import "fmt"

type Speaker interface{
	Speak()
}

type Dog struct{
	Name string
}

type Cat struct{
	Name string
}

type Robot struct{
	Name string
}

type Ghost struct{
	
}

func (d Dog) Speak() {
	fmt.Println("Woof!", d.Name)
}

func (c Cat) Speak() {
	fmt.Println("Meow!", c.Name)
}

func (r Robot) Speak() {
	fmt.Println("Beep Boop", r.Name)
}

func (g Ghost) Speak(){
	fmt.Println("Boooo!")
}

func main(){

	ensemble := []Speaker{
		Dog{Name: "Buddy"},
		Cat{Name: "Coco"},
		Robot{Name: "R2-D2"},
		Ghost{},
	}

	for _, memeber := range ensemble{
		memeber.Speak()
	}
}