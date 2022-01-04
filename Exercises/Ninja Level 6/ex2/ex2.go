package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := variadic(arr...)
	fmt.Println(sum)
}

func variadic(x ...int) int {
	totalSum := 0
	for _, v := range x {
		totalSum += v
	}
	return totalSum
}
