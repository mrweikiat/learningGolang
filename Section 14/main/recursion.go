package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)
}

func factorial(n int) int {
	// base case
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
