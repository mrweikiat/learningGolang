package main

import "fmt"

func main() {
	c := make(chan int)
	gen(c)
	for v := range c {
		fmt.Println(v)
	}
}

func gen(c chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}
