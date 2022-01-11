package main

import "fmt"

func main() {
	ans := rangeSum(1, 2, 3, 4, 5)
	fmt.Println(ans)

	// printing only the even numbers
	ans2 := even(rangeSum, 1, 2, 3, 4, 5, 6)
	fmt.Println(ans2)

	// printing only the odd numbers
	ans3 := odd(rangeSum, 1, 2, 3, 4, 5, 6)
	fmt.Println(ans3)
}

func rangeSum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, v ...int) int {
	var y []int
	for _, val := range v {
		if val%2 == 0 {
			y = append(y, val)
		}
	}
	return f(y...)
}

func odd(f func(x ...int) int, v ...int) int {
	var y []int
	for _, val := range v {
		if val%2 != 0 {
			y = append(y, val)
		}
	}
	return f(y...)
}
