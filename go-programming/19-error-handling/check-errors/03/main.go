package main

import (
	"io"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("sample.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("Sample")

	io.Copy(f, r)
}