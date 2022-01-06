package main

import "fmt"

func main() {
	a := 0
	foo(a)
	fmt.Println(a)
	bar(&a)
	fmt.Println(a)
}

func foo(x int) {
	fmt.Println(x)
	x = 43
	fmt.Println(x)
}

func bar(x *int) {
	fmt.Println(x)
	*x = 69
}
