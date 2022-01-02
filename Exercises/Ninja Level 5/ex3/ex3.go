package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println("This is a truck")
	fmt.Printf("The truck has %vdoors, is %v in color and it is %v that it is four wheeled\n", truck1.doors, truck1.color, truck1.fourWheel)
	fmt.Println("This is a sedan")
	fmt.Printf("The sedan has %vdoors, is %v in color and it is %v that it is a luxury car\n", sedan1.doors, sedan1.color, sedan1.luxury)

}
