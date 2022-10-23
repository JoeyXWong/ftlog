package bookstore

import (
	"errors"
)

type (

	// Book represents information about a book.
	Book struct {
		Id     int
		Title  string
		Author string
		Copies int
	}
)

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left to purchase")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	if book, found := catalog[id]; found {
		return book, nil
	}
	return Book{}, errors.New("Book for Id not found")
}
