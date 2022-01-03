package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	isExposed bool
}

// create a function that secret agent type can use
func (s secretAgent) kill() {
	fmt.Println("Executed killing of target,", s.first, s.last, "out")
}

func main() {
	jamesBond := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		isExposed: false,
	}
	fmt.Println(jamesBond)
	jamesBond.kill()
}
