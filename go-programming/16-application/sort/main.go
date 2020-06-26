package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{4, 3, 5, 1, 70, 33, 21, 59}
	y := []string{"n", "a", "i", "p", "q", "l", "b"}

	fmt.Println(x)
	fmt.Println(y)

	sort.Ints(x)
	fmt.Println(x)

	sort.Strings(y)
	fmt.Println(y)
}