package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "The Last Don",
		Author: "Mario Puzzo",
		Copies: 100,
	}
}

//Creating a function for our buy scenerio

func TestBuy(t *testing.T) {
	t.Parallel()
	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 47,
		},
		{
			Title:  "The Last Don",
			Author: "Mario Puzzo",
			Copies: 100,
		},
	}
	want := 46
	result, err := bookstore.Buy(books, "Splinter Cell")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	// The function returns the entire slice, so we need to
	// find our specific book to verify it was modified correctly
	var got int
	for _, b := range result {
		if b.Title == "Splinter Cell" {
			got = b.Copies
			break
		}
	}
	if want != got {
		t.Errorf("Error: %s, want %d, got %d", err, want, got)
	}
}

// TestStock tests the functionality of adding stock to an existing book
func TestStock(t *testing.T) {
	t.Parallel()
	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 47,
		},
	}
	want := 48
	result := bookstore.Stock(books, "Splinter Cell")
	//
	var got int
	for _, b := range result {
		if b.Title == "Splinter Cell" {
			got = b.Copies
			break
		}
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// TestBuyNoCopies tests the error scenario when trying to buy a book with no copies left
func TestBuyNoCopies(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 0,
		},
	}
	_, err := bookstore.Buy(books, "Splinter Cell")
	// The test passes if we get an error, fails if we don't.
	// This verifies our error handling logic works correctly.
	if err == nil {
		t.Error("expected error when no copies available, got nil")
	}
}

// TestBuyBookNotFound tests the error scenario when trying to buy a non-existent book
func TestBuyBookNotFound(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 47,
		},
	}
	_, err := bookstore.Buy(books, "Non-existent Book")
	if err == nil {
		t.Error("expected error when book not found, got nil")
	}
}

// TestGetAllBooks tests the retrieval of all books from inventory
func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Book 1",
			Author: "Author 1",
			Copies: 10,
		},
		{
			Title:  "Book 2",
			Author: "Author 2",
			Copies: 5,
		},
	}

	result := bookstore.GetAllBooks(books)

	// Ensuring we get back exactly what we put in,
	// no more and no less books than expected
	if len(result) != len(books) {
		t.Errorf("want %d books, got %d", len(books), len(result))
	}
}

// TestAddBook tests adding a new book to the inventory
func TestAddBook(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Existing Book",
			Author: "Existing Author",
			Copies: 10,
		},
	}

	newBook := bookstore.Book{
		Title:  "Purple Hibiscus",
		Author: "Chimamanda Ngozi Adichie",
		Copies: 15,
	}

	result := bookstore.AddBook(books, newBook)

	// Verify if we actually add one book
	if len(result) != 2 {
		t.Errorf("want %d books, got %d", 2, len(result))
	}

	// Length check alone is not sufficient, we need to
	// verify the specific book we added is actually in the result
	found := false
	for _, b := range result {
		if b.Title == "Purple Hibiscus" {
			found = true
			break
		}
	}

	if !found {
		t.Error("new book was not added to the slice")
	}
}
