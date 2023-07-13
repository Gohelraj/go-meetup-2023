package main

import (
	"fmt"
)

// Employee is an interface that defines the CalculateSalary method.
type Employee interface {
	CalculateSalary()
}

// Permanent struct represents a permanent employee.
type Permanent struct {
	empID       int
	basicSalary int
	pf          int
}

// Contract struct represents a contract employee.
type Contract struct {
	empID       int
	basicSalary int
}

// CalculateSalary calculates the salary for a permanent employee.
func (p Permanent) CalculateSalary() {
	fmt.Printf("Employee ID: %d, Salary: %d\n", p.empID, p.basicSalary+p.pf)
}

// CalculateSalary calculates the salary for a contract employee.
func (c Contract) CalculateSalary() {
	fmt.Printf("Employee ID: %d, Salary: %d\n", c.empID, c.basicSalary)
}

func main() {
	emp1 := Permanent{
		empID:       1,
		basicSalary: 15000,
		pf:          2000,
	}
	emp2 := Permanent{
		empID:       2,
		basicSalary: 10000,
		pf:          3000,
	}
	emp3 := Contract{
		empID:       3,
		basicSalary: 15500,
	}
	employees := []Employee{emp1, emp2, emp3}

	// Calculate and print the salary for each employee.
	for _, emp := range employees {
		emp.CalculateSalary()
	}
}
