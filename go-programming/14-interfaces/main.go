package main

import "fmt"

type person struct{
	firstName string
	lastName string
}

type batman struct{
	firstName string
}

type human interface{
	speak()
}

func (s person) speak() {
	fmt.Println("I can speak")
}

func (s batman) speak() {
	fmt.Println("I can speak")
}


func speak(h human) {
	fmt.Println("Speak from", h)
}

// Assertion
func speakAssertion(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I'm", h.(person).firstName)
	case batman:
		fmt.Println("I'm batman")
	}
}

func main() {
	p1 := person{
		firstName: "FirstName",
		lastName: "LastName",
	}

	p1.speak()

	speak(p1)

	speakAssertion(p1)

	p2 := batman{
		firstName: "Batman",
	}

	speakAssertion(p2)
}
