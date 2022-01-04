package main

import "fmt"

func main() {
	x := func(y int) int {
		y++
		return y
	}(42)
	fmt.Println(x)

}
