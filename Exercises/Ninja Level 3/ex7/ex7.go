package main

import "fmt"

func main() {
	x := 2
	// x := 4
	// x := 1
	y := 2
	if x == y {
		fmt.Println("x is equals to y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x  is smaller than y")
	}
}
