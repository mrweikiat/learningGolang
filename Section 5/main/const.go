package main

import "fmt"

// const is a keyword
const a = 123
const b = 69.999
const c = "I am James, James Bond"

// u can also declare like this in a package
const (
	xy int     = 43
	y  float32 = 1.2345
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T\n", a, b, c)
	fmt.Println(xy, y)
}
