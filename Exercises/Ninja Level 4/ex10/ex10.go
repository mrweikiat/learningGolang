package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`being evil`, `ice cream`, `sunsets`},
	}

	for k, v := range m {
		v = append(v, `dummy`)
		m[k] = v
	}

	for k, v := range m {
		v = append(v[:3], v[4:]...)
		m[k] = v
	}

	for i, v := range m {
		fmt.Println(i)
		for idx, elem := range v {
			fmt.Printf("\t%v %v\n", idx+1, elem)
		}
	}

}
