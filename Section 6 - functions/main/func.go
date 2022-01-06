package main

import "fmt"

func main() {
	foo()
}

// func (r receiver) identifier(parameters) (return(s)) {...}
func foo() {
	fmt.Println("I am foo")
	bar("John")
	dummy := woo("Some random gibberish")
	fmt.Println(dummy)
	a, b := mouse("James", "Bond")
	fmt.Println(a, ":", b)
}

func bar(name string) {
	fmt.Println("Hello,", name)
}

// how to take in a parameter and return a string
func woo(s string) string {
	return fmt.Sprintf("Hello from woo, this is %v", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, " says hello")
	b := true
	return a, b
}
