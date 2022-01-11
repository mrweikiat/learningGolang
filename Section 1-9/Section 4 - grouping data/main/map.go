package main

import "fmt"

func main() {
	fmt.Println()
	// map
	m := map[string]int{
		"James":  32,
		"Albert": 42,
		"Arthur": 69,
	}

	// printing the whole map
	fmt.Println(m)
	// printing the value using key
	fmt.Println(m["Arthur"])

	v, ok := m["James"]
	fmt.Println(v)  // print the value of the key, if no key then 0
	fmt.Println(ok) // this is to check if the key value exist
	z, okay := m["Jolene"]
	fmt.Println(z)
	fmt.Println(okay)

	// adding to a map
	m["Michael"] = 999
	fmt.Println(m)

	// print key and value pair iterative
	for k, v := range m {
		fmt.Println(k, v)
	}

	// deleting an element from map
	delete(m, "James")
	delete(m, "Fake name")
	fmt.Println(m)

}
