package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello my name is", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
	person1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	person1.speak()
}
