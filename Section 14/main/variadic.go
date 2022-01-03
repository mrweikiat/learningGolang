package main

import "fmt"

func main() {
	alpha(1, 2, 3)

}

func alpha(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		fmt.Println("for item in index position,", i, "we are adding ", v)
		sum = sum + v
	}
	fmt.Println("Total sum is ", sum)
}
