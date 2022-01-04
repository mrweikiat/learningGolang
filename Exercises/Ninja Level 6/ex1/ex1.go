package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
	y, name := bar()
	fmt.Println(name, "is", y, "years old")
}

func foo() int {
	return 69
}

func bar() (int, string) {
	return 42, "James Bond"
}
