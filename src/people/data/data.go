package data

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Person struct {
	FirstName string `csv:"first_name"`
	LastName  string `csv:"last_name"`
}

func LoadPeopleFromCSV(filePath string) ([]Person, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var people []Person
	if err := gocsv.UnmarshalFile(file, &people); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CSV: %w", err)
	}
	return people, nil
}