package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println("Sort by age:")
	for _, v := range users {
		fmt.Printf("\t%v %v age is %v\n", v.First, v.Last, v.Age)
	}

	sort.Sort(ByName(users))
	fmt.Println("Sorting by last name now:")
	for _, v := range users {
		fmt.Printf("\t%v %v\n", v.First, v.Last)
	}

	sort.Sort(BySaying(users))
	fmt.Println("Sorting by sayings now:")
	for _, v := range users {
		fmt.Printf("\t%v %v said:%v\n", v.First, v.Last, v.Sayings)
	}

}

type ByAge []user

func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByName []user

func (a ByName) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BySaying []user

func (a BySaying) Less(i, j int) bool { return len(a[i].Sayings) < len(a[j].Sayings) }
func (a BySaying) Len() int           { return len(a) }
func (a BySaying) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
