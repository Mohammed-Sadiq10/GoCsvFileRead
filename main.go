package main

import (
	csvRead "github.com/Mohammed-Sadiq10/GoCsvFileRead/CSV"
	jsonOutput "github.com/Mohammed-Sadiq10/GoCsvFileRead/JSON"
	"fmt"
	"log"
)

func main() {

	employees, err := csvRead.ReadCSVFile("employees.csv")
	if err != nil {
		log.Fatal("Error reading CSV data:", err)
	}

	log.Println("Processed Employee Data:")
	for _, emp := range employees {
		fmt.Printf("\nID: %d, Name: %s %s, Age: %d, Salary: %.2f, IsActive: %v\n", emp.ID, emp.FirstName, emp.LastName, emp.Age, emp.Salary, emp.IsActive)
	}

	err = jsonOutput.WriteJSONFile("output.json", employees)
	if err != nil {
		log.Fatal("Error writing JSON data:", err)
	}

	log.Println("JSON data has been written to output.json")
}