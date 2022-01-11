package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 69
	}()

	fmt.Println(<-c)

	//buffered channel
	anotherC := make(chan int, 1)
	anotherC <- 123
	fmt.Println(<-anotherC)
}
