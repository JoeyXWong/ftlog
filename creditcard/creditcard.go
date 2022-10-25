package creditcard

import "fmt"

type (
	card struct {
		number string
	}
)

func New(number string) (card, error) {
	if number == "" {
		return card{}, fmt.Errorf("cannot create card with empty number")
	}
	c := card{number: number}
	return c, nil
}

func (c card) GetNumber() string {
	return c.number
}
