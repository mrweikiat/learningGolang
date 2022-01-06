package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var wg1 sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(3)

	go foo()
	go bar()
	go doSmth()

	wg.Wait()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("Hi im foo")
	wg.Done()
}

func bar() {
	fmt.Println("Hi im bar")
	wg.Done()
}

func doSmth() {
	fmt.Println("I am doing something!!!")
	wg.Done()
}
