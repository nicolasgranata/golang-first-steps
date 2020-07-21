package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	f, err := os.Create("error-log.txt")

	if err != nil {
		fmt.Println("Something happened from fmt", err)
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("file.txt")
	
	if err != nil {
		log.Println("Something happened from log", err)
		fmt.Println("Check the errorlog.txt for errors")
	}

	defer f2.Close()
}