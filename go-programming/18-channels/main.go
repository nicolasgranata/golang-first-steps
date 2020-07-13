package main

import "fmt"

func main() {
	// Pass of value
	c := make(chan int)

	go func () {
		c <- 25
	}()

	fmt.Println(<-c)

	// Buffer
	d := make(chan int, 1)

	d <- 26

	fmt.Println(<-d)

	// Buffer
	e := make(chan int, 2)

	e <- 27
	e <- 28

	fmt.Println(<-e)
	fmt.Println(<-e)
}