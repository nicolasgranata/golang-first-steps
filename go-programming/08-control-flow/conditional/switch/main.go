package main

import "fmt"

func main() {
	x := "My name"

	switch x {
	case "My name":
		fmt.Println("This is my name")
	case "Not my name":
		fmt.Println("This is not my name")
	default:
		fmt.Println("What's your name?")
	}	
}