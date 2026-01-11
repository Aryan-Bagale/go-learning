package main

import "fmt"

type Player struct {
	Name string
	Level int
	Alive bool
}

func (p *Player) Levelup(){
	p.Level++
	fmt.Println(p.Name, "leveled up!")
}

func (p * Player) KnockOut(){
	p.Alive = false
	fmt.Println("Oh no! Hero was knocked out.")
}


func main(){
	p1 := Player{
		Name: "Hero",
		Level: 10,
		Alive : true,
	}

	fmt.Println("Start Level:", p1.Level)

	p1.Levelup()

	fmt.Println("Start Level:", p1.Level)

	p1.KnockOut()

	fmt.Println("Is Alive:", p1.Alive)

}
