package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Player struct {
	Username string
	Level    int
	Score    int
}

func main() {
	// 1. LOAD current state
	data, err := os.ReadFile("savegame.json")
	if err != nil {
		panic(err)
	}

	var p Player
	if err := json.Unmarshal(data, &p); err != nil {
		panic(err)
	}
	fmt.Printf("Current Level: %d\n", p.Level)

	// 2. MODIFY (Level Up!)
	p.Level = p.Level + 1
	fmt.Println("Leveling up...")

	// 3. SAVE new state
	newData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	
	err = os.WriteFile("savegame.json", newData, 0644)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Saved! New Level: %d\n", p.Level)
}