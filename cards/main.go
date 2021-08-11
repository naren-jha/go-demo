package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"

	// card := "Ace of Spades" // declaring and assiging a new variable
	// fmt.Println(card)

	// card = "Five of Diamonds" // reassiging an existing variable
	// fmt.Println(card)

	card := newCard()
	fmt.Println(card)

	fmt.Println(getCountry()) // comes from file2.go
	// run as "go run main.go file2.go"
}

func newCard() string {
	return "Five of Diamonds"
}
