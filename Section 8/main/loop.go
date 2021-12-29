package main

import (
	"fmt"
)

func main() {

	// a typical for loop in Go
	for i := 0; i < 3; i++ {
		fmt.Println("hello")
	}

	// nested loops in Go
	// we can use same instances of i unlike in java and kotlin
	for i := 0; i < 3; i++ {
		fmt.Println("Im in the outer loop of", i)
		for i := 0; i < 2; i++ {
			fmt.Println("Im in the inner loop of", i)
		}
	}

	// for statement with a single condition
	x := 2
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// break and continue in for loops
	temp := 80
	for {
		if temp > 100 {
			break // continue also works like break
		}
		if temp%2 == 0 {
			fmt.Println(temp)
		}
		temp++
	}

}
