package main

import "fmt"

func main(){
	var team = []string{"Alice","Bob"}

	fmt.Println("Start:", team)

	team = append(team, "Charlie")
	team = append(team,"David")

	fmt.Println("Updated:", team)
	fmt.Println("Team Size:", len(team))
}