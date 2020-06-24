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

	sliceInt := []int{1,2,3,4,5}
	foo5(sliceInt...)

	// Defer
	defer foo6()
	bar()

	p := person {
		name: "My Name",
	}

	p.walk()

	// Anonymous func
	func() {
		fmt.Println("Anonymous func")
	}()

	func(x int) {
		fmt.Println("Anonymous func", x)
	}(30)

	// Func expression
	fex := func() {
		fmt.Println("Func expression")
	}

	fex2 := func(x int) {
		fmt.Println("Func expression passing parameters", x)
	}

	fex()
	fex2(30)
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

func foo6() {
	fmt.Println("Foo6")
}

func bar() {
	fmt.Println("Bar")
}

// Attaching a Method to a struct
type person struct {
	name string
}

func (p person) walk() {
	fmt.Println("I'm walking")
}