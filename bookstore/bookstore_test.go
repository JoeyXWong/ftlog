package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBookStore(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}
}

func TestBuyBook(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "book1",
		Author: "author1",
		Copies: 1,
	}

	want := 0
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatalf("error in the buying step with valid copies")
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "book1",
		Author: "author1",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)
	if err == nil {
		t.Fatalf("should throw an error for zero copies purchase")
	}

}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "book1"},
		{Title: "book2"},
	}

	want := []bookstore.Book{
		{Title: "book1"},
		{Title: "book2"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Id: 1},
		3: {Id: 3},
	}

	want := bookstore.Book{Id: 3}

	got, err := bookstore.GetBook(catalog, 3)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(got, want))
	}
}

func TestGetInvalidBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Id: 1},
	}

	_, err := bookstore.GetBook(catalog, 3)
	if err == nil {
		t.Errorf("Expected error for missing Id lookup")
	}
}
