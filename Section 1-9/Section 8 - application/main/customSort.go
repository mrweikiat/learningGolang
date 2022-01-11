package main

import (
	"fmt"
	"sort"
)

type person1 struct {
	first string
	age   int
}

type ByAge []person1

func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByName []person1

func (ba ByName) Less(i, j int) bool { return ba[i].first < ba[j].first }
func (ba ByName) Len() int           { return len(ba) }
func (ba ByName) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }

func (p person1) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

func main() {
	p1 := person1{"James", 32}
	p2 := person1{"Moneypenny", 27}
	p3 := person1{"Q", 64}
	p4 := person1{"M", 56}

	people := []person1{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
