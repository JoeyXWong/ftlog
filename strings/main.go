package main

import "fmt"

type (
	Customer struct {
		Name  string
		Email string
	}
)

func main() {
	var title string
	var copies int
	title = "For the love of go"
	copies = 99
	fmt.Println(title)
	fmt.Println(copies)

	firstEdition := 1
	secondEdition := 2
	fmt.Println(firstEdition)
	fmt.Println(secondEdition)

	// inStock := true
	// royaltyPercentage := 12.5
	// isDiscounted := true

}
