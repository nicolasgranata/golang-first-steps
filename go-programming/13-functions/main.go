package main

import "fmt"

func main() {
	foo()

	foo2("Foo2")

	s := foo3("Foo3")
	fmt.Println(s)

	a, b := foo4("Hello", "Bye")
	fmt.Println(a, b)

	// Variadic parameter
	foo5(2,3,4,5)
}

func foo() {
	fmt.Println("Foo")
}

func foo2(s string) {
	fmt.Println(s)
}

func foo3(s string) string {
	return s
}

func foo4(s string, s2 string) (string, bool) {
	return s, true
}

// Variadic parameter
func foo5(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}