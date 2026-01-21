package main

import "fmt"

type Character struct {
	name string
	hp int
}

func takeDamage(c *Character, damage int){
	c.hp = c.hp - damage
}

func heal( c *Character){
	c.hp = 100
}

func main(){
	hero := Character{name: "Knight", hp: 100}

	takeDamage(&hero, 20)

	fmt.Println("HP after damage:", hero.hp)

	heal(&hero)

	fmt.Println("HP after heal:", hero.hp)
}