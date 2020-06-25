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

	fmt.Println(foo7()())
	fmt.Printf("%T\n", foo7())

	ii := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenNumbers := even(sum, ii...)
	fmt.Println("even numbers", evenNumbers)

	factorial := factorial(4)
	fmt.Println(factorial)
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

// Return a function
func foo7() func() int {
	return func() int {
		return 123
	}
}


// Callback
func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi [] int
	for _,v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n  * factorial(n - 1)
}