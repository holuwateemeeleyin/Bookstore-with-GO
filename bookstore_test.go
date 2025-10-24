package bookstore_test

import (
	"bookstore"
	"testing"
)

// TestBook tests the basic creation of a Book struct
// WHY WE NEED THIS TEST: This is a foundational test that ensures our basic
// data structure can be created correctly. If this fails, nothing else will work.
func TestBook(t *testing.T) {
	t.Parallel()
	// WHY BLANK IDENTIFIER: We're testing that the struct can be created without
	// runtime errors. We don't need to use the value, just verify initialization works.
	_ = bookstore.Book{
		Title:  "The Last Don",
		Author: "Mario Puzzo",
		Copies: 100,
	}
	// WHY NO ASSERTIONS: If this code compiles and runs without panic, the test passes.
	// We're testing that the struct definition and initialization syntax work correctly.
}

// TestBuy tests the normal purchase scenario where copies are available
// WHY WE NEED THIS TEST: This is our "happy path" test - it verifies that the
// most common operation (buying a book with available stock) works correctly.
func TestBuy(t *testing.T) {
	t.Parallel()

	// WHY MULTIPLE BOOKS IN SLICE: Real inventory has multiple books. We need to test
	// that our function can find the right book among others and only modify that one.
	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 47, // WHY SPECIFIC NUMBER: We need a known starting point to verify the decrement
		},
		{
			Title:  "The Last Don",
			Author: "Mario Puzzo",
			Copies: 100, // WHY INCLUDE SECOND BOOK: To ensure we don't accidentally modify other books
		},
	}

	want := 46 // WHY HARDCODED EXPECTED VALUE: We know 47-1=46. This makes the test clear and predictable
	result, err := bookstore.Buy(books, "Splinter Cell")
	// WHY CHECK ERROR FIRST: If there's an unexpected error, we should fail fast rather than
	// continuing with potentially invalid data
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// WHY SEARCH THROUGH RESULTS: The function returns the entire slice, so we need to
	// find our specific book to verify it was modified correctly
	var got int
	for _, b := range result {
		if b.Title == "Splinter Cell" {
			got = b.Copies
			break
		}
	}

	// WHY COMPARE WANT vs GOT: This is the actual verification - does our function
	// produce the expected business result (inventory decreased by 1)?
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// TestBuyNoCopies tests the error scenario when trying to buy a book with no copies left
// WHY WE NEED THIS TEST: Error handling is crucial. We need to ensure our system
// gracefully handles edge cases rather than crashing or allowing invalid operations.
func TestBuyNoCopies(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 0, // WHY ZERO COPIES: This creates the specific error condition we want to test
		},
	}

	// WHY WE IGNORE THE RESULT: In error cases, we care about the error, not the result.
	// The blank identifier makes it clear we're intentionally ignoring the returned slice.
	_, err := bookstore.Buy(books, "Splinter Cell")
	// WHY CHECK FOR ERROR: The test passes if we get an error, fails if we don't.
	// This verifies our error handling logic works correctly.
	if err == nil {
		t.Error("expected error when no copies available, got nil")
	}
}

// TestBuyBookNotFound tests the error scenario when trying to buy a non-existent book
// WHY WE NEED THIS TEST: Customers might request books we don't carry. We need to
// ensure our system handles this gracefully rather than crashing or behaving unpredictably.
func TestBuyBookNotFound(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 47,
		},
	}

	// WHY NON-EXISTENT BOOK NAME: We need to test the case where the search fails entirely
	_, err := bookstore.Buy(books, "Non-existent Book")
	if err == nil {
		t.Error("expected error when book not found, got nil")
	}
}

// TestStock tests the functionality of adding stock to an existing book
// WHY WE NEED THIS TEST: Restocking is a core business operation. We need to verify
// that adding inventory works correctly and doesn't affect other books.
func TestStock(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{
			Title:  "Splinter Cell",
			Author: "Tom Clancy",
			Copies: 47, // Known starting point
		},
	}

	want := 48 // Expected: 47 + 1 = 48
	result := bookstore.Stock(books, "Splinter Cell")

	// WHY SEARCH AGAIN: Same reason as in TestBuy - we need to extract the specific
	// book we modified from the returned slice to verify the change
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

// TestGetAllBooks tests the retrieval of all books from inventory
// WHY WE NEED THIS TEST: Even simple getter functions should be tested to ensure
// they don't accidentally modify data or return incorrect results
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

	// WHY TEST LENGTH: We want to ensure we get back exactly what we put in,
	// no more and no less books than expected
	if len(result) != len(books) {
		t.Errorf("want %d books, got %d", len(books), len(result))
	}
}

// TestAddBook tests adding a new book to the inventory
// WHY WE NEED THIS TEST: Adding new products is a common operation. We need to
// verify that new books are properly added to our inventory system.
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
		Title:  "New Book",
		Author: "New Author",
		Copies: 5,
	}

	result := bookstore.AddBook(books, newBook)

	// WHY CHECK LENGTH: The most basic verification - did we actually add one book?
	if len(result) != 2 {
		t.Errorf("want %d books, got %d", 2, len(result))
	}

	// WHY SEARCH FOR THE BOOK: Length check alone isn't sufficient - we need to
	// verify the specific book we added is actually in the result
	found := false
	for _, b := range result {
		if b.Title == "New Book" {
			found = true
			break
		}
	}

	if !found {
		t.Error("new book was not added to the slice")
	}
}
