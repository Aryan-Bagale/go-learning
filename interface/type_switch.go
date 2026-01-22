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
        Ghost{},
        Robot{Name: "R2-D2"},
    }

    for _, member := range ensemble{

        member.Speak()

        switch v:= member.(type){
        case Dog:
            fmt.Println("  -> treat for " + v.Name)
        case Cat:
            fmt.Println("  ->" + v.Name + " is sleeping")
        case Ghost:
            fmt.Println("  -> Run away!")
        case Robot:
            fmt.Println("  -> Charging battery...")
        }
    }
}