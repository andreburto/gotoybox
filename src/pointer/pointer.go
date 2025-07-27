package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Hello() {
	fmt.Printf("Hello, my name is %s.\n", p.FullName())
}

func setName(name *[]Person) int{
	// var ppl []People
	*name = append(*name, Person{FirstName: "John", LastName: "Doe"})
	*name = append(*name, Person{FirstName: "Jane", LastName: "Smith"})
	return len(*name)
}

func main() {
	var ppl []Person
	count := setName(&ppl)
	fmt.Println("Number of people:", count)

	for i := 0; i < len(ppl); i++ {
		ppl[i].Hello()
	}
}