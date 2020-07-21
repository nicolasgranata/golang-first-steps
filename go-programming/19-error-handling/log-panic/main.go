package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("file.txt")

	if err != nil {
		log.Panicln(err)
	}
}