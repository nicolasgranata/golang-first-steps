package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Age int
}

func main() {
	p1 := person {
		First: "James",
		Last: "Bond",
		Age: 32,
	}

	fmt.Println(p1)

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}