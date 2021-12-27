package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z) // 0 , "", false
	// no values assign to var -> will be default values
}
