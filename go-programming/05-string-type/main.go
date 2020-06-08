package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		// UTF-8 Code
		fmt.Printf("%#U", s[i])
	}

	for i := 0; i < len(s); i++ {
		// UTF-8 HEX
		fmt.Printf("%#x", s[i])
	}
}