package main

import "fmt"

func main() {

	favSport := "soccer"

	switch favSport {
	case "soccer":
		fmt.Println("favourite sport is soccer")
	case "volleyball":
		fmt.Println("favourite sport is volleyball")
	default:
		fmt.Println("i love all sports")
	}
}
