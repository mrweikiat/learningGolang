package main

import "fmt"

func main() {
	proxy := []int{1, 2, 3, 4, 5}
	ans := sum(proxy...)
	fmt.Println("The sum is", ans)
	// u can also pass in zero arguments
	noAns := sum()
	fmt.Println("The sum is", noAns)

	// passing multiple arguments into a func
	name := stringAndSum("James Bond", 1, 2, 3, 4, 5)
	fmt.Println(name)

}

func sum(x ...int) int {
	fmt.Println(x)
	sum := 0

	for _, v := range x {
		sum += v
	}
	return sum
}

func stringAndSum(s string, x ...int) int {

	sum := 0

	for _, v := range x {
		sum += v
	}

	fmt.Println(s, "is", sum, "years old")
	return sum
}
