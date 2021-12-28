package main

import "fmt"

func main() {
	x := 42
	y := 69
	s := x == y
	fmt.Println(s)
	s = x >= y
	fmt.Println(s)
	s = x <= y
	fmt.Println(s)
	s = x != y
	fmt.Println(s)
	s = x < y
	fmt.Println(s)
	s = x > y
	fmt.Println(s)
}
