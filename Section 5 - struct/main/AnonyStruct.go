package main

import "fmt"

func main() {

	// creating an instance object within a funct
	personOne := struct {
		first string
		last  string
	}{
		first: "John",
		last:  "Doe",
	}

	fmt.Println(personOne.first)
	fmt.Println(personOne.last)

}
