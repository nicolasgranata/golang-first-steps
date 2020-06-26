package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "Pass123"
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
	fmt.Println(bs)

	loginPassword := "Pass123"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("Error login")
	}

	fmt.Println("Login")
}