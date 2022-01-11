package main

import "fmt"

func main() {
	//send
	c := make(chan int)

	go fo(c)
	// no goroutine for printing func
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Program Exited")
}

// func to send
func fo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // closing the channel
}
