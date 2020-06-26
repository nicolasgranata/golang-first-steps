package main

import "fmt"

func main() {
	a := 30
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // get the value stored at an address

	*b = 31
	fmt.Println(a)

	x := 0
	foo(&x)
	fmt.Println(x)
}


func foo(y *int) {
	fmt.Println(*y)
	*y = 29
	fmt.Println(*y)
}