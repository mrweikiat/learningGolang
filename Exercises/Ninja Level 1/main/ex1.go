package main

import "fmt"

// i used global var here but should be short declaration -> should be in the main func
var x = 42
var y = "James Bond"
var z bool = true

func main() {
	// this is to print multiple types into a SINGLE print statement

	x1 := 42
	y1 := "James Bond"
	z1 := true

	// different ways of printing
	s := fmt.Sprint(x, " ", y, " ", z)
	fmt.Println(s)
	fmt.Println(x, y, z)
	fmt.Println(x1, y1, z1)

	// printing each type individually is easy
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
