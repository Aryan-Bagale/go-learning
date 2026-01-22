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

func (d Dog) Speak(){
	fmt.Println("Woof! I am", d.Name)
}

func (c Cat) Speak(){
	fmt.Println("Meow!", c.Name)
}

func main(){

	var myAnimal Speaker

	myAnimal = Dog{Name:"buddy"}

	myAnimal.Speak()

	myAnimal = Cat{Name:"Coco"}

	myAnimal.Speak()
}
