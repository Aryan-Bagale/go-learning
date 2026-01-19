package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Player struct{
	Username string
	Level int
	Score int
}

func main(){
	p := Player{Username: "DragonSlayer", Level: 55, Score: 1000}

	data, err := json.Marshal(p)
	if err != nil{
		panic(err)
	}

	os.WriteFile("savegame.json", data, 0644)
	fmt.Println("Game saved:", string(data))

}