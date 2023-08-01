package csvutil

import (
	model "github.com/Mohammed-Sadiq10/GoCsvFileRead/Model"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)


func CsvFileReading(csvFilePath string) ([]model.Employee, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []model.Employee
	for i, row := range csvData {
		if i == 0 {
			continue 
		}

		id, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[3])
		hireDate, _ := time.Parse("2006-01-02", row[8])
		salary, _ := strconv.ParseFloat(row[9], 64)
		isActive, _ := strconv.ParseBool(row[12])

		employee := model.Employee{
			ID:         id,
			FirstName:  row[1],
			LastName:   row[2],
			Age:        age,
			Gender:     row[4],
			Email:      row[5],
			Phone:      row[6],
			Address:    row[7],
			HireDate:   hireDate,
			Salary:     salary,
			Department: row[10],
			Title:      row[11],
			IsActive:   isActive,
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
