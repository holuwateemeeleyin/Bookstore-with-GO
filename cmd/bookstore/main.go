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
			Copies: 15, // Multiple books to show slice functionality
		},
	}

	// Showing the starting state so we can see changes
	fmt.Println("Initial bookstore inventory:")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// Trying to buy to show the system in action
	fmt.Println("\n--- Processing a purchase ---")
	var err error
	books, err = bookstore.Buy(books, "Things Fall Apart")

	// Checking if there is an error
	if err != nil {
		fmt.Printf("Error buying book: %s\n", err)
	} else {
		fmt.Println("After buying 'Things Fall Apart':")
		for _, book := range books {
			fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
		}
	}

	// Showing the complete inventory management cycle
	fmt.Println("\n--- Restocking a book ---")
	books = bookstore.Stock(books, "The Last Don")
	fmt.Println("After stocking 'The Last Don':")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// The inventory should grow over time
	fmt.Println("\n--- Adding a new book to inventory ---")
	newBook := bookstore.Book{
		Title:  "New Book",
		Author: "New Author",
		Copies: 10, // New book with initial stock
	}
	books = bookstore.AddBook(books, newBook)
	fmt.Println("After adding a new book:")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// The end state and demonstrating the GetAllBooks functionality
	fmt.Println("\n--- Final Inventory Summary ---")
	allBooks := bookstore.GetAllBooks(books)
	fmt.Printf("Total books in inventory: %d\n", len(allBooks))
}
