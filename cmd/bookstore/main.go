package main

import (
	"bookstore"
	"fmt"
)

// We can also define it using a shortcut :=
func main() {
	//Bookstore inventory
	books := []bookstore.Book{
		{
			Title:  "Things Fall Apart",
			Author: "Chinua Achebe",
			Copies: 20,
		},
		{
			Title:  "The Last Don",
			Author: "Mario Puzzo",
			Copies: 15,
		},
	}

	// Showing the starting state so we can see changes
	fmt.Println("Initial bookstore inventory:")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// Buying a new book
	fmt.Println("\n--- Processing a purchase ---")
	//var err error
	books = bookstore.Buy(books, "Things Fall Apart")
	fmt.Println("After buying 'Things Fall Apart':")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// Restocking
	fmt.Println("\n--- Restocking a book ---")
	books = bookstore.Stock(books, "The Last Don")
	fmt.Println("After stocking 'The Last Don':")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}
}
