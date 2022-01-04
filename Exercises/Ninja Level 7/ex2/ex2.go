package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).age = 123
	(*p).first = "Alan"
	(*p).last = "Walker"
}

func main() {
	p1 := person{
		"James",
		"Bond",
		23,
	}
	fmt.Println("Hi I am", p1.first, p1.last, "and I am", p1.age, "years old")

	changeMe(&p1)
	fmt.Println("Hi I am", p1.first, p1.last, "and I am", p1.age, "years old")

}
