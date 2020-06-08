package main

import "fmt"

const a int = 30
const b float64 = 30.25
const c string = "Hello World"

const (
	d = 31
	e = 31.25
	f= "Hello"
)

func main() {
	fmt.Println(a, b ,c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d, e, f)

	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}