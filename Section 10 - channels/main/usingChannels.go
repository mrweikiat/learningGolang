package main

import "fmt"

func main() {
	//send
	c := make(chan int)

	go foo(c)
	// no goroutine for printing func
	bar(c)

	fmt.Println("Program Exited")
}

// func to send
func foo(c chan<- int) {
	c <- 123
}

// func to rev
func bar(c <-chan int) {
	fmt.Println(<-c)
}
