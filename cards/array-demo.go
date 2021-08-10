package main

import "fmt"

/*
	Array - is like array in Java
	Slice - like ArrayList in Java
	Slice or Array are containers of same data type

	[]datatype{} - declaring a Slice
*/

func main() {
	cards := []string{"Ace of Spades", newCard()}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for index, card := range cards {
		fmt.Println(index, card)
	}
	/*
		0 Ace of Spades
		1 Five of Diamonds
		2 Six of Spades
	*/

	for index := range cards {
		fmt.Println(index)        // 0 1 2
		fmt.Println(cards[index]) // {entries}
	}
}

func newCard() string {
	return "Five of Diamonds"
}
