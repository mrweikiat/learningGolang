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

	s1 := fmt.Sprintf("%v %v's favourite icecream is %v", person1.firstName, person1.lastName, person1.favouriteIceCream)
	s2 := fmt.Sprintf("%v %v's favourite icecream is %v", person2.firstName, person2.lastName, person2.favouriteIceCream)

	m := map[string]string{
		"Doe":  s1,
		"Bond": s2,
	}

	for _, v := range m {
		fmt.Println(v)
	}

}
