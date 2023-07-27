package model

import "time"

type Employee struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	HireDate   time.Time `json:"hire_date"`
	Salary     float64   `json:"salary"`
	Department string    `json:"department"`
	Title      string    `json:"title"`
	IsActive   bool      `json:"is_active"`
}