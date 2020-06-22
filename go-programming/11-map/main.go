package main

import "fmt"

func main() {
	m := map[int]string{
		1: "banana",
		2: "apple",
		3: "orange",
	}

	fmt.Println(m)

	fmt.Println(m[1])

	fmt.Println(m[2])

	fmt.Println(m[4])

	// Check if exist
	v, ok := m[4]

	fmt.Println(v)
	fmt.Println(ok)

	// Check if exist
	if v, ok := m[3]; ok {
		fmt.Println(v)
	}

	// add new entry
	m[4] = "pineapple"

	for k, v := range m {
		fmt.Println(k, v)
	}

	// remove entry
	delete(m, 2)

	for k, v := range m {
		fmt.Println(k, v)
	}
}