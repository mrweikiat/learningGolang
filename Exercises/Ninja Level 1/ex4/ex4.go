package main

import "fmt"

type wildCard int

var x wildCard

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
