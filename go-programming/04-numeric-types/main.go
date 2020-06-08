package main

import "fmt"

var a int
var b float64

func main() {
	x := 29
	y := 29.2928
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	a = 29
	b = 29.2928
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
