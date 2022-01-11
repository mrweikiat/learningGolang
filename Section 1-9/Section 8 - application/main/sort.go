package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{4, 3, 2, 1, 5}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}
