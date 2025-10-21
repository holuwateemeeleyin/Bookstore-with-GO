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

// Buy processes the purchase of a book by decreasing its copy count by 1
// We pass in the book we want to buy, and once the purchase is processed
// our new book is returned to us.
func Buy(books []Book, title string) ([]Book, error) {
	//We have multiple books in our inventory, so we need to search
	// through all of them to find the specific book the customer wants to buy
	for i, b := range books {
		if b.Title == title {
			//We can't sell what we don't have
			if b.Copies == 0 {
				return books, errors.New("no copies left")
			}
			//Once we sold a copy, we update our inventory by reducing the number
			books[i].Copies--
			return books, nil
		}
	}
	return books, errors.New("books not found")
}

// Stock increases the copy count of a specific book by 1 when a new book arrives
func Stock(books []Book, title string) []Book {
	//we need to find book to restock and increment the count
	for i, b := range books {
		if b.Title == title {
			//Each book once found and added, it increases our stocks
			books[i].Copies++
			return books
		}
	}
	//return slice if we don't find any book to increment
	return books
}

// Get all books function
func GetAllBooks(books []Book) []Book {
	return books
}

// Add a new book entirely that is not in our stock
func AddBook(books []Book, book Book) []Book {
	return append(books, book)
}
