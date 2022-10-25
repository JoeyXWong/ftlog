package bookstore

import (
	"errors"
	"fmt"
)

type (

	// Book represents information about a book.
	Book struct {
		Id              int
		Title           string
		Author          string
		Copies          int
		PriceCents      int
		DiscountPrecent int
		category        Category
	}

	Catalog map[int]*Book

	Category int
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left to purchase")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	var books []Book
	for _, b := range c {
		books = append(books, *b)
	}
	return books
}

func (c Catalog) GetBook(id int) (Book, error) {
	if book, found := c[id]; found {
		return *book, nil
	}
	return Book{}, fmt.Errorf("Book for Id:%d not found", id)
}

func (b Book) GetNetPriceCents() float64 {
	netPrice := float64(b.Copies*b.PriceCents) * (float64(b.DiscountPrecent) / 100.00)
	return netPrice
}

func (b *Book) SetPriceForBook(priceCents int) error {
	if priceCents < 0 {
		return fmt.Errorf("cannot set negative price: %d", priceCents)
	}
	b.PriceCents = priceCents
	return nil
}

func (b *Book) SetCatageory(category Category) error {
	if _, ok := validCategory[category]; !ok {
		return fmt.Errorf("setting an invalid category.")
	}

	b.category = category
	return nil
}

func (b Book) GetCategory() Category {
	return b.category
}
