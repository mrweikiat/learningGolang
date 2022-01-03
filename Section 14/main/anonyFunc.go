package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("hello i am anony")
	}()

	func(x int) {
		fmt.Println("Hello my number is", x)
	}(999)

	// func expression, assign value to variable
	f := func(x int) {
		fmt.Println("this is a func with a number called", x)
	}
	f(69)

	// returning a func from a func
	s1 := returnFunc()
	fmt.Println(s1)
	xx := anotherReturnFunc()
	fmt.Println(xx) // returning func() int hence its the memory address
	i := xx()
	fmt.Println(i)
}

func returnFunc() string {
	s := "I will be back"
	return s
}

func anotherReturnFunc() func() int {
	return func() int {
		return 123
	}
}
