package main

import "fmt"

func main() {
	// use keyword defer to run the function at the end of this func
	defer forDefer()
	forDefer2()
}

func forDefer() {
	fmt.Println("this is for testing")
}

func forDefer2() {
	fmt.Println("this is another testing")
}
