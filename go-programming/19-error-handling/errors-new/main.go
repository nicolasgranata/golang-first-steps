package main

import (
	"errors"
	"log"
)

func main() {
	_, err := printPositiveNumber(- 1)

	if err != nil {
		log.Fatalln(err)
	}
}

func printPositiveNumber(number int) (int, error) {
	if number < 0 {
		return 0, errors.New("Number should be positive")
	}

	return number, nil
}