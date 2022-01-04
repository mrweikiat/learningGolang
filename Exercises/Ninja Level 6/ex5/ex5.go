package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	breath float64
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) area() float64 {
	return s.length * s.breath
}

func (c CIRCLE) area() float64 {
	r2 := c.radius * c.radius
	return math.Pi * r2
}

type SHAPE interface {
	area() float64
}

func main() {
	sq := SQUARE{
		length: 10,
		breath: 5,
	}

	//areaSq := sq.area()
	//fmt.Println(areaSq)

	cir := CIRCLE{
		radius: 10,
	}
	//areaCircle := cir.area()
	//fmt.Println(areaCircle)

	info(sq)
	info(cir)
}

func info(s SHAPE) {
	fmt.Println(s.area())
}
