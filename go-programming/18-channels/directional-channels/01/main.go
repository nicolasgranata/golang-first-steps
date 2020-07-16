package main

import "fmt"

func main() {
	cr := make(<-chan int) // receive
	cs := make(chan <- int) // send

	fmt.Println("-----")
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}