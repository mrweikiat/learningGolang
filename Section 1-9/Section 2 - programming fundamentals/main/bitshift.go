package main

import "fmt"

const (
	_ = iota
	//kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	//gb = mb * 1024
	gb = 1 << (iota * 10)
)

func main() {
	two := 2
	fmt.Printf("%d\t%b\n", two, two)

	// use the operator << to shift up 1 bit
	shift := two << 1
	fmt.Printf("%d\t%b\n", shift, shift)

	fmt.Println(kb, mb, gb)

}
