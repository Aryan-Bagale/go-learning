package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Player struct {
	Username string
	Level int
	Score int
}

func main(){
	data, err := os.ReadFile("savegame.json")
	if err != nil{
			panic(err)
	}

	var loadedPlayer Player
	err = json.Unmarshal(data, &loadedPlayer)
	if err != nil{
			panic(err)
	}
	fmt.Printf("Loaded: %+v\n", loadedPlayer)
	fmt.Println("Score is:", loadedPlayer.Score)
}