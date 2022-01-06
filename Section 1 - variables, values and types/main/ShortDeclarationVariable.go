// := is used for short declaration variable
package main

import "fmt"

var y = 44 // use var when u need to create a global variable
var z int  // create a default int variable, int by default is 0

func main() {
	// use this if u need to create variables used only in the func call
	x := 69
	fmt.Println(x)
	// u can reassign the value using =
	x = 123
	fmt.Println(x)
	var y = 44
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("Im in foo")
}
