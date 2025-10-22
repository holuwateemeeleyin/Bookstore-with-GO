package bookstore_test

import (
	"bookstore"
	"testing"
)

// Creating a function for our buy scenerio
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
	result := bookstore.Buy(books, "Splinter Cell")
	// The function returns the entire slice, so we need to
	// find our specific book to verify it was modified correctly
	var got int
	for _, b := range result {
		if b.Title == "Splinter Cell" {
			got = b.Copies
		}
	}
	if want != got {
		t.Errorf("Want %d, got %d", want, got)
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
		}
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
