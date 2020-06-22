package main

import "fmt"

func main() {
	// x:= type{values} 
	// composite literal
	x := []int{4, 5, 7, 8, 30}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	x = append(x, 31, 32)
	fmt.Println(x)

	y := []int{40, 41, 42}
	x = append(x, y...)
	fmt.Println(x)

	// delete
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// slice with make
	a := make([]int, 10, 20)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a[0] = 30
	a[9] = 39
	fmt.Println(a)

	// due to the capacity (20)
	a = append(a, 40)
	fmt.Println(a)

	// multi-dimensional slice
	b := []string{"One", "Two", "Three"}
	c := []string{"Four, Five, Six"}
	d := [][]string{b, c}
	fmt.Println(d)
}