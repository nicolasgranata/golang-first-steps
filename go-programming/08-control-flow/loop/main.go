package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++{
		fmt.Println(i)
	} 

	// break and continue
	x := 0
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
}