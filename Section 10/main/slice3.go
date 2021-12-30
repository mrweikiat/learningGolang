package main

import "fmt"

func main() {

	x := make([]int, 10, 100)
	fmt.Println(x)
	x = append(x, 99, 100)
	fmt.Println(x)

	z := [][]int{{1, 2}, {3, 4}}
	fmt.Println(z)

}
