package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	anotherC := make(chan int, 1)
	anotherC <- 123
	fmt.Println(<-anotherC)
}
