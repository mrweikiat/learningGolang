package main

import "fmt"

func main() {
	ans := []int{1, 2, 3, 4, 5, 6}
	i := countEvenOnly(rangeSum, ans...)
	fmt.Println(i)

}

func rangeSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func countEvenOnly(f func(x ...int) int, v ...int) int {
	var y []int
	for _, val := range v {
		if val%2 == 0 {
			y = append(y, val)
		}
	}
	return f(y...)
}