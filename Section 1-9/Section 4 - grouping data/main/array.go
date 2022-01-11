package main

import "fmt"

func main() {
	var x [5]int
	var y [6]int // x is not equal to y !!!
	// length of arr is part of the arr's type
	fmt.Println(x)
	x[3] = 999
	fmt.Println(x)
	fmt.Println(y)
}
