package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}

	x = append(x[:0], x[5:]...)
	x = append(x, 42, 43, 44, 45, 46)
	fmt.Println(x)
	x = append(x[:0], x[5:]...)
	x = append(x, 47, 48, 49, 50, 51)
	fmt.Println(x)
	x = append(x[:0], x[5:]...)
	x = append(x, 44, 45, 46, 47, 48)
	fmt.Println(x)
	x = append(x[:0], x[5:]...)
	x = append(x, 43, 44, 45, 46, 47)
	fmt.Println(x)

}
