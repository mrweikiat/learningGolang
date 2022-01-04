package main

import "fmt"

func main() {
	x := returnFunc()
	i := x()
	fmt.Println(i)
}

func returnFunc() func() int {
	return func() int {
		return 123
	}
}
