package grasp_patterns

import (
	"fmt"
)

// BooksModule handles book-related operations
type BooksModule struct {
	books  []Book
	nextID int
}

type Book struct {
	ID     int
	Title  string
	Author string
}

func (bm *BooksModule) AddBook(title, author string) {
	bm.nextID++
	book := Book{ID: bm.nextID, Title: title, Author: author}
	bm.books = append(bm.books, book)
}

// PatronsModule handles patron-related operations
type PatronsModule struct {
	patrons []Patron
}

type Patron struct {
	ID   int
	Name string
}

func (pm *PatronsModule) RegisterPatron(name string) {
	patron := Patron{ID: len(pm.patrons) + 1, Name: name}
	pm.patrons = append(pm.patrons, patron)
}

// LibraryModule coordinates interactions between different modules
type LibraryModule struct {
	booksModule   BooksModule
	patronsModule PatronsModule
}

func (lm *LibraryModule) RegisterPatron(name string) {
	lm.patronsModule.RegisterPatron(name)
}

func (lm *LibraryModule) AddBook(title, author string) {
	lm.booksModule.AddBook(title, author)
}

func main() {
	library := LibraryModule{}

	library.RegisterPatron("Alice")
	library.RegisterPatron("Bob")

	library.AddBook("The Great Gatsby", "F. Scott Fitzgerald")
	library.AddBook("To Kill a Mockingbird", "Harper Lee")

	fmt.Println("Library Management System:")
	fmt.Println("Registered Patrons:", library.patronsModule.patrons)
	fmt.Println("Available Books:", library.booksModule.books)
}
