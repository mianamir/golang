package solid_principles

type Book struct {
	Title  string
	Author string
	Price  float64
}

type Inventory struct {
	Books []Book
}

func (inv *Inventory) AddBook(book Book) {
	inv.Books = append(inv.Books, book)
}

func (inv *Inventory) RemoveBook(book Book) {
	for i, b := range inv.Books {
		if b == book {
			inv.Books = append(inv.Books[:i], inv.Books[i+1:]...)
			break
		}
	}
}
