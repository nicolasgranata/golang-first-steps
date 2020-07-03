package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// run this code with go run -race main.go

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i :=0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GOroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}