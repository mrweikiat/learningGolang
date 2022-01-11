package main

import "fmt"

// this is like OOP where we can create instances
type person struct {
	firstName string
	lastName  string
	age       int
}

// this is like inheritance
// in Go its called embedding another type with struct
type athlete struct {
	person
	topSpeed int
}

func main() {
	person1 := person{
		firstName: "John",
		lastName:  "Doe",
	}

	person2 := person{
		firstName: "weikiat",
		lastName:  "goh",
	}

	fmt.Println(person1.lastName)
	fmt.Println(person1.firstName)
	fmt.Println(person2.lastName)
	fmt.Println(person2.firstName)
	fmt.Println(person1 == person2)
	fmt.Println(person1 == person1)

	// now we can create an instance of athlete
	sprinter := athlete{
		person: person{
			firstName: "Usain",
			lastName:  "Bolt",
		},
		topSpeed: 99,
	}
	fmt.Printf("%v %v, has a top speed of %v\n", sprinter.firstName, sprinter.person.lastName, sprinter.topSpeed)
}
