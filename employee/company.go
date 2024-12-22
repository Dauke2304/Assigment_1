package employee

import "fmt"

type Company struct {
	employees map[string]Employee
}

// AddEmployee adds an employee to the company
func (c *Company) AddEmployee(emp Employee) {
	if c.employees == nil {
		c.employees = make(map[string]Employee)
	}
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.employees[e.Name] = emp
	case PartTimeEmployee:
		c.employees[e.Name] = emp
	}
}

// ListEmployees lists all employees in the company
func (c Company) ListEmployees() {
	fmt.Println("Listing all employees:")
	for _, emp := range c.employees {
		fmt.Println(emp.GetDetails())
	}
}
