package grasp_patterns

import (
	"fmt"
)

// Patron defines the interface for patrons
type Patron interface {
	HasExceededBorrowLimit() bool
}

// Student implements Patron
type Student struct {
	// ...
}

// Faculty implements Patron
type Faculty struct {
	// ...
}

func (s *Student) HasExceededBorrowLimit() bool {
	// Logic to check if a student has exceeded the borrow limit
	return false
}

func (f *Faculty) HasExceededBorrowLimit() bool {
	// Logic to check if a faculty member has exceeded the borrow limit
	return true
}

func main() {
	student := &Student{}
	faculty := &Faculty{}
	patrons := []Patron{student, faculty}
	for _, patron := range patrons {
		if !patron.HasExceededBorrowLimit() {
			fmt.Println("Allow borrowing for patron")
		} else {
			fmt.Println("Borrow limit exceeded for patron")
		}
	}
}
