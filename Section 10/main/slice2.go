package main

import "fmt"

func main() {
	// adding elem to a slice
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	x = append(x, 42, 69) // akin to adding element into the slice
	fmt.Println(x)

	// delete elem from a slice
	y := []int{42, 69, 365}
	y = append(y[:1], y[3:]...) // specify the start and end of the range u want to delete
	fmt.Println(y)
	// this example we remove 69 and 365

	z := [][]int{{1, 2}, {3, 4}}
	fmt.Println(z)
}
