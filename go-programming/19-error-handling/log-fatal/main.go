package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("file.txt")

	if err != nil {
		// Defer functions don't run
		log.Fatalln(err)
	}
}