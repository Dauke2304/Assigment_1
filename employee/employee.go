package employee

import (
	"fmt"
	"strconv"
)

// Employee interface
type Employee interface {
	GetDetails() string
}

// FullTimeEmployee struct
type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

// PartTimeEmployee struct
type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

// GetDetails for FullTimeEmployee
func (f FullTimeEmployee) GetDetails() string {
	return "FullTimeEmployee: ID: " + strconv.FormatUint(f.ID, 10) + ", Name: " + f.Name + ", Salary: " + strconv.Itoa(int(f.Salary)) + " Tenge"
}

// GetDetails for PartTimeEmployee
func (p PartTimeEmployee) GetDetails() string {
	return "PartTimeEmployee: ID: " + strconv.FormatUint(p.ID, 10) + ", Name: " + p.Name + ", Hourly Rate: " + strconv.FormatUint(p.HourlyRate, 10) + " Tenge, Hours Worked: " + fmt.Sprintf("%.2f", p.HoursWorked)
}
