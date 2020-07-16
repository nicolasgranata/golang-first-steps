package main

import "fmt"

func main() {
	c := make(chan int)
	
	// send
	go foo(c)

	// receive
	// pull until the channel is closed
	for v:= range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func foo(c chan<- int) {
	for i:=0; i < 100; i++ {
		c <- i
	}
	close(c)
}