package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Printf("%T\n", x)

	a := 7
	b := 8
	fmt.Println(a == b)
	fmt.Println(a != b)
}
