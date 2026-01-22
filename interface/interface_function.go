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

func (d Dog) Speak() {
	fmt.Println("Woof!", d.Name)
}

func (c Cat) Speak() {
	fmt.Println("Meow!", c.Name)
}

func (r Robot) Speak() {
	fmt.Println("Beep Boop", r.Name)
}

func MakeItTalk(s Speaker) {
	fmt.Println("--- Mic Check ---")
	s.Speak()
}

func main(){

	dog := Dog{Name:"Buddy"}
	cat := Cat{Name:"Coco"}
	robot := Robot{Name:"R2-D2"}

	MakeItTalk(dog)
	MakeItTalk(cat)
	MakeItTalk(robot)

}