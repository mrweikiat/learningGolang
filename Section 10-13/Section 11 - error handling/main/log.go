package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("troller.txt")
	if err != nil {
		fmt.Println("error happened...can't read file")
	}

	f2, error := os.Open("another-fake-file.txt")
	if error != nil {
		log.Println("error!!!", error)
	}
	defer f2.Close()
	fmt.Println("check log file in dir")

	_, error2 := os.Open("another-fake-file.txt")
	if error2 != nil {
		log.Fatalln("error!!!", error)
	}

}
