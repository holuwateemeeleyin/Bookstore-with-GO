package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}
type Customer struct {
	Name    string
	Address string
	Email   string
}
type Author struct {
	Name  string
	Email string
}

// Buy enables the purchase of a book from our store
// We pass in the book we want to buy, and once the purchase is processed
// our new book is returned to us.
func Buy(b Book) (Book, error) {
	//if there is no book to sell
	if b.Copies == 0 {
		return b, errors.New("no copies left")
	}
	//Given thst books purchased reduce in number...
	//b.Copies = b.Copies - 1
	// or more succinctly
	b.Copies--
	//return the new book
	return b, nil
}

func Stock(b Book) Book {
	b.Copies++
	return b
}
