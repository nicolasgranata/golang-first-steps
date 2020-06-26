package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type ByAge []Person

func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	people := []Person{
		{"Juan", 31},
		{"Ana", 42},
		{"Pedro", 17},
		{"Marta", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}