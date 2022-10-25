package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	catalog := bookstore.Catalog{
		1: {Id: 1, Title: "book1"},
		2: {Id: 2, Title: "book2"},
	}

	want := []bookstore.Book{
		{Id: 1, Title: "book1"},
		{Id: 2, Title: "book2"},
	}

	got := catalog.GetAllBooks()

	sort.Slice(got, func(i, j int) bool {
		return got[i].Id < got[j].Id
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Id: 1},
		3: {Id: 3},
	}

	want := bookstore.Book{Id: 3}

	got, err := catalog.GetBook(3)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(got, want))
	}
}

func TestGetInvalidBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Id: 1},
	}

	_, err := catalog.GetBook(3)
	if err == nil {
		t.Errorf("Expected error for missing Id lookup")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Book{Id: 1, Copies: 5, PriceCents: 52, DiscountPrecent: 50}

	got := catalog.GetNetPriceCents()

	want := 130.0

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSetPriceCentsValid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "title1", PriceCents: 100}

	want := 200

	err := b.SetPriceForBook(200)
	if err != nil {
		t.Fatalf(err.Error())
	}

	got := b.PriceCents

	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{}

	err := b.SetPriceForBook(-100)
	if err == nil {
		t.Errorf("Should receive error for invalid book id retrieval")
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{}
	err := b.SetCatageory(bookstore.Category(100))
	if err == nil {
		t.Errorf("Should receive error for setting an invalid category")
	}
}

func TestSetCategoryValid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{}

	categories := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}

	for _, category := range categories {
		err := b.SetCatageory(category)

		if err != nil {
			t.Fatalf(err.Error())
		}

		got := b.GetCategory()
		want := category

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}
}
