package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`being evil`, `ice cream`, `sunsets`},
	}

	for i, v := range m {
		fmt.Println(i)
		for idx, elem := range v {
			fmt.Printf("\t%d %v\n", idx+1, elem)
		}
	}

}
