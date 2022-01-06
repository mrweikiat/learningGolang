package main

import "fmt"

func main() {
	if true {
		fmt.Println("hello")
	}
	if !false {
		fmt.Println("hello")
	}

	x := 2

	// if elseif else
	for x < 4 {
		if x == 2 {
			fmt.Println("x is", x)
		} else if x != 2 {
			fmt.Println("x is not equals to 2, x is now", x)
		} else {

		}
		x++
	}

	y := 5

	// fallthrough can  let u evaluate other cases, NOT RECOMMENDED
	switch {
	case (y == 2):
		fmt.Println("y is equals to", y)
		y++
		//fallthrough
	case true:
		fmt.Println("this is always true")
		//fallthrough
	case (y == 3):
		fmt.Println("y is equal to", y)
		//fallthrough
	default:
		fmt.Println("this is default case")
	}

	// switch statements with expression
	z := 14
	switch z {
	case 14:
		fmt.Println("z is", z)
	case 999:
		fmt.Println("Should not print")

	}
	// && and
	// || or
}
