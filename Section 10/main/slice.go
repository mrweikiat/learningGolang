package main

import "fmt"

func main() {

	// composite literal
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))
	// slice can put tgt values of the same type
	fmt.Println(x[0])
	for i, v := range x {
		fmt.Println(i, v) // v is the element
	}
	// its like java
	// for (int element : arr) {
	// println(element)
	// }

	// we only print position 1 to 3 inclusive
	fmt.Println(x[1:3])

	// iterate the slice w/o range clause
	for i := 0; i <= 3; i++ {
		fmt.Println(i, x[i])
	}
}
