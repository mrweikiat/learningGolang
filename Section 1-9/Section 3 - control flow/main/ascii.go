package main

import "fmt"

func main() {
	fmt.Println("helloe")
	x := 'a'
	character := int(x)
	fmt.Println(character)

	for i := 33; i < 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
