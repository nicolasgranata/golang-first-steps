package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}