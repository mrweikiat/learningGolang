package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 123
	b := 999
	fmt.Println(a == b)
	fmt.Println(a != b)
}
