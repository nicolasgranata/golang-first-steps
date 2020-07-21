package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	_, err := os.Open("sample.txt")

	if err != nil {
		fmt.Println("Something happened from fmt", err)
		log.Println("Something happened from log", err)
	}
}