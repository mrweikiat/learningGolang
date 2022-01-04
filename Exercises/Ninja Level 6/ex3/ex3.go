package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("hello this is foo")
}

func bar() {
	fmt.Println("hello this is bar")
}
