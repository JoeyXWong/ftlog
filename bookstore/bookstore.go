package bookstore

import (
	"errors"
	"fmt"
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

func GetAllBooks(catalog map[int]Book) []Book {
	var books []Book
	for _, b := range catalog {
		books = append(books, b)
	}
	return books
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	if book, found := catalog[id]; found {
		return book, nil
	}
	return Book{}, fmt.Errorf("Book for Id:%d not found", id)
}
