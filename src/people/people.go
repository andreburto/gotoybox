package main

import (
	"fmt"

	"data"
	"display"
)

func main() {
	people, err := data.LoadPeopleFromCSV("people.csv")
	if err != nil {
		fmt.Printf("Failed to load people: %v\n", err)
		return
	}

	for _, person := range people {
		display.Hello(person.FirstName)
	}	
}
