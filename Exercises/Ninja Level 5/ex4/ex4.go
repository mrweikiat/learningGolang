package main

import "fmt"

func main() {
	// creating an instance object within a funct
	john := struct {
		height int
		weight int
		name   string
	}{
		height: 180,
		weight: 70,
		name:   "John Doe",
	}

	fmt.Printf("Hello, my name is %v. I am %vcm tall and %vkg heavy.", john.name, john.height, john.weight)

}
