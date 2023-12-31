package main

import (
	csvRead "github.com/Mohammed-Sadiq10/GoCsvFileRead/CSV"
	jsonOutput "github.com/Mohammed-Sadiq10/GoCsvFileRead/JSON"
	"log"
)

func main() {

	employees, err := csvRead.CsvFileReading("employees.csv")
	if err != nil {
		log.Fatal("Error reading CSV data:", err)
	}

	log.Println("Processed Employee Data:")
	for _, emp := range employees {
		log.Printf("ID: %d, Name: %s %s, Email: %s, Phone: %s, Age: %d, Salary: %.2f, IsActive: %v\n", emp.ID, emp.FirstName, emp.LastName, emp.Email, emp.Phone, emp.Age, emp.Salary, emp.IsActive)
	}

	err = jsonOutput.JsonFileOutput("jsonOutput.json", employees)
	if err != nil {
		log.Fatal("Error writing JSON data:", err)
	}

	log.Println("JSON data has been written to jsonOutput.json")
}