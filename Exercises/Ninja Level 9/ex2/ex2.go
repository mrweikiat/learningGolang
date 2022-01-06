package main

import "fmt"

type person struct {
	first  string
	last   string
	age    int
	speech string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello I am speaking blah")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"james",
		"bond",
		32,
		"hello",
	}

	saySomething(&p1)
	p1.speak()

}
