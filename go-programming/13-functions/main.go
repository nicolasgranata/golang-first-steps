package main

import "fmt"

func main() {
	foo()

	foo2("Foo2")

	s := foo3("Foo3")
	fmt.Println(s)

	a, b := foo4("Hello", "Bye")
	fmt.Println(a, b)
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