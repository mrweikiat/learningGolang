package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	var wait sync.WaitGroup

	const gs = 100
	wait.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//time.Sleep(time.Second)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
