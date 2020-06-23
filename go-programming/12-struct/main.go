package main

import "fmt"

type person struct{
	firstName string
	lastName string
	job
}

type job struct{
	description string
}

func main() {
	p1 := person { 
		firstName: "Michael",
		lastName: "Jordan",
		job: job {
			description: "Basketball player",
		},
	}

	p2 := person {
		firstName: "Lebron",
		lastName: "James",
		job: job {
			description: "Basketball player",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p2.firstName)

	fmt.Println(p1.firstName, p1.description)

	// Anonymous structs
	p3 := struct {
		name string
	} {
		name: "This is my name",
	}

	fmt.Println(p3.name)
}