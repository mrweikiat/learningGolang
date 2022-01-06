package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	var wait sync.WaitGroup
	var mutex sync.Mutex

	const gs = 100
	wait.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			mutex.Lock()

			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v

			mutex.Unlock()
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
