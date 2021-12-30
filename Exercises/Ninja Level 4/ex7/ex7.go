package main

import "fmt"

func main() {

	x := [][]string{{"James", "Bond", "Shaken", "Not stirred"},
		{"Miss", "Moneypenny", "Hello", "James"}}

	for i := range x {
		y := x[i]
		fmt.Printf("%v ", i)
		for elem := range y {
			fmt.Printf("%v ", y[elem])
		}
		fmt.Println()
	}
}
