package main

import "fmt"

func main() {
	x := 2

	if x == 0 {
		fmt.Println("zero")
	}

	if x != 0 {
		fmt.Println("Not zero")
	}

	if  x % 2 == 0 {
		fmt.Println("It's even")
	} else if x % 2 != 0 {
		fmt.Println("It's odd")
	} else {
		fmt.Println("?")
	}		
}