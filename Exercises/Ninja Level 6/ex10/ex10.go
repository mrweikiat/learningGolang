package main

import "fmt"

func main() {
	a := increment()
	ans := a()
	fmt.Println(ans)
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
