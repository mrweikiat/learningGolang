package main

import "fmt"

var fourTwo = 42                                       // this is of type int
var text = "This is a dummy text, dont think too hard" // this is of type string

// gonna create my own type, named ryan and of type int
type ryan int

var b ryan

func main() {

	fmt.Println(fourTwo)
	fmt.Println(text)
	fmt.Printf("%T\n", fourTwo) // this is to print out the type of the var
	fmt.Printf("%T\n", text)

	s := fmt.Sprintf(string(fourTwo))
	fmt.Println(s)
	fourTwo = 69
	fmt.Println(fourTwo)

	b = 999
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// type casting
	fourTwo = int(b) // if we type cast b into int, we can do this
	fmt.Println(fourTwo)
	
}
