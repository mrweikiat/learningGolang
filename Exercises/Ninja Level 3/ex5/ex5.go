package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		x := i % 4
		fmt.Println(x)
	}
}
