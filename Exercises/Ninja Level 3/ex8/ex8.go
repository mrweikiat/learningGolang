package main

import "fmt"

func main() {

	x := 42
	y := 69

	switch {
	case x == 42:
		fmt.Println("x is equals to 42")
		fallthrough
	case y == 69:
		fmt.Println("y is equals to 69")
	default:
		fmt.Println("x is 42 and y is 69")
	}

}
