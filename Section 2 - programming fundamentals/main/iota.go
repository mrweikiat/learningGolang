package main

import "fmt"

// you can do this when you need something to increment by ++
const (
	aa = iota + 1
	bb
	cc
)

func main() {
	fmt.Println(aa, bb, cc)
	fmt.Printf("%T %T %T\n", aa, bb, cc)
}
