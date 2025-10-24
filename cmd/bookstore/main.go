package main

import (
	"bookstore"
	"fmt"
)

func main() {
	// WHY INITIALIZE WITH DATA: A real bookstore starts with existing inventory.
	// This demonstrates our system working with actual data from the beginning.
	books := []bookstore.Book{
		{
			Title:  "Things Fall Apart",
			Author: "Chinua Achebe",
			Copies: 20, // Realistic starting inventory
		},
		{
			Title:  "The Last Don",
			Author: "Mario Puzzo",
			Copies: 15, // Multiple books show slice functionality
		},
	}

	// WHY DISPLAY INITIAL INVENTORY: Shows the starting state so we can see changes
	fmt.Println("Initial bookstore inventory:")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// WHY DEMONSTRATE BUYING: This is the core business operation - shows the system in action
	fmt.Println("\n--- Processing a purchase ---")
	var err error
	books, err = bookstore.Buy(books, "Things Fall Apart")
	// WHY CHECK ERROR: In real applications, we must handle errors gracefully
	if err != nil {
		fmt.Printf("Error buying book: %s\n", err)
	} else {
		fmt.Println("After buying 'Things Fall Apart':")
		for _, book := range books {
			fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
		}
	}

	// WHY DEMONSTRATE STOCKING: Shows the complete inventory management cycle
	fmt.Println("\n--- Restocking a book ---")
	books = bookstore.Stock(books, "The Last Don")
	fmt.Println("After stocking 'The Last Don':")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Copies: %d\n", book.Title, book.Author, book.Copies)
	}

	// WHY DEMONSTRATE ADDING NEW BOOK: Shows how the inventory grows over time
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

	// WHY FINAL SUMMARY: Shows the end state and demonstrates the GetAllBooks functionality
	fmt.Println("\n--- Final Inventory Summary ---")
	allBooks := bookstore.GetAllBooks(books)
	fmt.Printf("Total books in inventory: %d\n", len(allBooks))
	// WHY PRINT TOTAL COUNT: Gives a quick overview of inventory size
}
