package grasp_patterns

import (
	"fmt"
)

// PatronServices defines the interface for handling patron-related operations
type PatronServices interface {
	PatronCanBorrowBook(patron Patron, book Item) bool
}

// LibraryServices handles library-related operations
type LibraryServices struct {
	patronServices PatronServices
	books          []Book
	nextBookID     int
}

type Book struct {
	ID     int
	Title  string
	Author string
}

type Patron struct {
	ID   int
	Name string
}

type Item interface {
	// ...
}

func (ls *LibraryServices) AddBook(title, author string) {
	ls.nextBookID++
	book := Book{ID: ls.nextBookID, Title: title, Author: author}
	ls.books = append(ls.books, book)
}

func (ls *LibraryServices) CanBorrowBook(patron Patron, book Item) bool {
	return ls.patronServices.PatronCanBorrowBook(patron, book)
}

// ConcretePatronServices implements PatronServices
type ConcretePatronServices struct {
	// ...
}

func (cps *ConcretePatronServices) PatronCanBorrowBook(patron Patron, book Item) bool {
	// Logic to determine if the patron can borrow the book
	return true
}

func main() {
	library := LibraryServices{
		patronServices: &ConcretePatronServices{},
	}

	library.AddBook("The Great Gatsby", "F. Scott Fitzgerald")
	library.AddBook("To Kill a Mockingbird", "Harper Lee")

	patron := Patron{ID: 1, Name: "Alice"}
	book := Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}

	canBorrow := library.CanBorrowBook(patron, book)
	fmt.Printf("Can %s borrow %s? %v\n", patron.Name, book.Title, canBorrow)
}
