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
	b := bookstore.Book{
		Title:  "Splinter Cell",
		Author: "Tom Clancy",
		Copies: 47,
	}
	want := 46
	result, err := bookstore.Buy(b)
	got := result.Copies
	if want != got {
		t.Errorf("Error: %s, want %d, got %d", err, want, got)
	}
}

func TestStock(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Splinter Cell",
		Author: "Tom Clancy",
		Copies: 47,
	}
	want := 48
	result := bookstore.Stock(b)
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
