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
	s := `{"First":"James", "Last":"Bond","Age":32}`
	bs := []byte(s)

	fmt.Println(s)

	var p1 person

	err := json.Unmarshal(bs, &p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1)
}