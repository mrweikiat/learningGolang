package main

import "fmt"

type monkey struct {
	first string
	last  string
}

type gorilla struct {
	monkey
	isStrong bool
}

func (g gorilla) smash() {
	fmt.Println(g.first, g.last, "LOVES TO SMASH")
}

func (m monkey) throwBananas() {
	fmt.Println(m.first, m.last, "is throwing bananas NOW")
}

type homo interface {
	throwBananas()
}

func main() {
	monke := gorilla{
		monkey: monkey{
			"King",
			"Kong",
		},
		isStrong: true,
	}
	smallBoy := gorilla{
		monkey: monkey{
			"White",
			"Baboon",
		},
		isStrong: false,
	}

	normalMonkey := monkey{
		first: "Monkey D",
		last:  "Luffy",
	}

	fmt.Println(monke)

	monke.smash()
	smallBoy.smash()

	homo.throwBananas(monke)
	homo.throwBananas(smallBoy)
	homo.throwBananas(normalMonkey)
	
	smallBoy.throwBananas()
	normalMonkey.throwBananas()
}
