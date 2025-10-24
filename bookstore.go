package bookstore

import "errors"

// Book struct represents a book in our bookstore inventory
// WHY WE NEED THIS: In real-world applications, we need a structured way to represent
// our data. A struct allows us to group related data together (title, author, copies)
// making it easier to manage and pass around related information as a single unit.
type Book struct {
	Title  string // WHY: Store the book title to identify the book
	Author string // WHY: Store author information for display and search
	Copies int    // WHY: Track inventory count to know when we need to reorder
}

// Buy processes the purchase of a book by decreasing its copy count by 1
// WHY WE NEED THIS FUNCTION: In a real bookstore, when a customer buys a book,
// we need to reduce our inventory count. This function encapsulates that business logic.
func Buy(books []Book, title string) ([]Book, error) {
	// WHY WE ITERATE: We have multiple books in our inventory, so we need to search
	// through all of them to find the specific book the customer wants to buy
	for i, b := range books {
		if b.Title == title {
			// WHY WE CHECK COPIES: Business requirement - we can't sell what we don't have.
			// This prevents negative inventory counts which would cause data integrity issues.
			if b.Copies == 0 {
				return books, errors.New("no copies left")
			}
			// WHY WE DECREMENT: This is the core business logic - one copy sold means
			// one less copy in inventory. This updates our stock levels accurately.
			books[i].Copies--
			return books, nil
		}
	}
	// WHY WE RETURN ERROR IF NOT FOUND: The customer might request a book we don't carry.
	// We need to handle this gracefully rather than silently failing.
	return books, errors.New("book not found")
}

// Stock increases the copy count of a specific book by 1
// WHY WE NEED THIS FUNCTION: When new shipments arrive or books are returned,
// we need to increase our inventory counts. This handles the restocking process.
func Stock(books []Book, title string) []Book {
	// WHY WE SEARCH BY TITLE: We need to find the exact book to restock among
	// potentially hundreds of titles in our inventory
	for i, b := range books {
		if b.Title == title {
			// WHY WE INCREMENT: Each new copy added to inventory increases our stock level.
			// This ensures we have accurate counts for future sales.
			books[i].Copies++
			return books
		}
	}
	// WHY WE RETURN ORIGINAL SLICE: If we don't find the book, we can't stock it.
	// Returning the unchanged slice is better than creating an error for this simple operation.
	return books
}

// GetAllBooks returns the current slice of all books in inventory
// WHY WE NEED THIS FUNCTION: Other parts of our application (like UI or reporting)
// need to see the entire inventory. This provides a clean interface to access all books
// without exposing the internal slice implementation directly.
func GetAllBooks(books []Book) []Book {
	// WHY SIMPLE RETURN: This acts as a "getter" function - it provides controlled
	// access to our data while allowing us to add validation or logging later if needed.
	return books
}

// AddBook adds a new book to the inventory slice
// WHY WE NEED THIS FUNCTION: When we start carrying a new book title, we need to
// add it to our inventory system. This handles the onboarding of new products.
func AddBook(books []Book, book Book) []Book {
	// WHY WE USE APPEND: Slices in Go are dynamic, but we need to properly add
	// new elements. Append handles the underlying array management for us.
	// WHY WE RETURN THE NEW SLICE: Append might return a new slice with a different
	// memory address, so we need to return it so the caller has the updated version.
	return append(books, book)
}
