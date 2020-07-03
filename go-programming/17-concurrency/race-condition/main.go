package main

import ( 
	"fmt"
	"runtime"
	"sync"
)

// run this code with go run -race main.go

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i :=0; i < gs; i++ {
		go func(){
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GOroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}