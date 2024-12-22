// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	library "github.com/Dauke2304/Assignment_1/Library"
// )

// func main() {
// 	lib := library.NewLibrary()
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		fmt.Println("\n1. Add")
// 		fmt.Println("2. Borrow")
// 		fmt.Println("3. Return")
// 		fmt.Println("4. List")
// 		fmt.Println("5. Exit")
// 		fmt.Print("Choose an option: ")

// 		if !scanner.Scan() {
// 			break
// 		}
// 		option := strings.ToLower(scanner.Text())

// 		switch option {
// 		case "add":
// 			fmt.Print("Enter book ID: ")
// 			scanner.Scan()
// 			id := scanner.Text()

// 			fmt.Print("Enter book title: ")
// 			scanner.Scan()
// 			title := scanner.Text()

// 			fmt.Print("Enter book author: ")
// 			scanner.Scan()
// 			author := scanner.Text()

// 			err := lib.AddBook(library.Book{ID: id, Title: title, Author: author})
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 			} else {
// 				fmt.Println("Book added successfully.")
// 			}

// 		case "borrow":
// 			fmt.Print("Enter book ID to borrow: ")
// 			scanner.Scan()
// 			id := scanner.Text()

// 			err := lib.BorrowBook(id)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 			} else {
// 				fmt.Println("Book borrowed successfully.")
// 			}

// 		case "return":
// 			fmt.Print("Enter book ID to return: ")
// 			scanner.Scan()
// 			id := scanner.Text()

// 			err := lib.ReturnBook(id)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 			} else {
// 				fmt.Println("Book returned successfully.")
// 			}

// 		case "list":
// 			lib.ListBooks()

// 		case "exit":
// 			fmt.Println("Exiting program.")
// 			return

// 		default:
// 			fmt.Println("Invalid option. Please try again.")
// 		}
// 	}

// }

// package main

// import (
// 	"fmt"

// 	shapes "github.com/Dauke2304/Assignment_1/Shapes"
// )

// func main() {
// 	shapesList := []shapes.Shape{
// 		shapes.Rectangle{Length: 5, Width: 3},
// 		shapes.Circle{Radius: 4},
// 		shapes.Square{Length: 6},
// 		shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
// 	}

// 	for i, shape := range shapesList {
// 		fmt.Printf("Shape %d Details:\n", i+1)
// 		shapes.PrintShapeDetails(shape)
// 		fmt.Println()
// 	}
// }

package main

import "github.com/Dauke2304/Assigment_1/employee"

func main() {
	// Initialize company
	company := employee.Company{}

	// Create employees
	fullTimeEmp := employee.FullTimeEmployee{ID: 1, Name: "Alice", Salary: 300000}
	partTimeEmp := employee.PartTimeEmployee{ID: 2, Name: "Bob", HourlyRate: 5000, HoursWorked: 20.5}

	// Add employees to the company
	company.AddEmployee(fullTimeEmp)
	company.AddEmployee(partTimeEmp)

	// List all employees
	company.ListEmployees()
}
