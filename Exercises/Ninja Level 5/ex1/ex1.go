package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteIceCream string
}

func main() {

	person1 := person{
		firstName:         "John",
		lastName:          "Doe",
		favouriteIceCream: "Vanilla softserve",
	}
	person2 := person{
		firstName:         "James",
		lastName:          "Bond",
		favouriteIceCream: "Peppermint",
	}

	fmt.Printf("%v %v's favourite icecream is %v\n", person1.firstName, person1.lastName, person1.favouriteIceCream)
	fmt.Printf("%v %v's favourite icecream is %v\n", person2.firstName, person2.lastName, person2.favouriteIceCream)

}
