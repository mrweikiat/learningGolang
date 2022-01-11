package main

import "fmt"

func main() {
	// this channel can only send values into the channel
	c := make(chan<- int, 2)
	c <- 123
	c <- 42

	//cr := make(<-chan int)
	//cs := make(chan<- int)
	
	// so u cannot pull out the val from c to print
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)

}
