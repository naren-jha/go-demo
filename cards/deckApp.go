package main

func main() {
	cards := newDeck() // function in deck.go
	// cards.print()      // function in deck.go
	/*
		0 Ace of Spades
		1 Two of Spades
		2 Three of Spades
		3 Four of Spades
		4 Ace of Diamonds
		5 Two of Diamonds
		6 Three of Diamonds
		7 Four of Diamonds
		8 Ace of Hearts
		9 Two of Hearts
		10 Three of Hearts
		11 Four of Hearts
		12 Ace of Clubs
		13 Two of Clubs
		14 Three of Clubs
		15 Four of Clubs
	*/

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	/*
		0 Ace of Spades
		1 Two of Spades
		2 Three of Spades
		3 Four of Spades
		4 Ace of Diamonds
		0 Two of Diamonds
		1 Three of Diamonds
		2 Four of Diamonds
		3 Ace of Hearts
		4 Two of Hearts
		5 Three of Hearts
		6 Four of Hearts
		7 Ace of Clubs
		8 Two of Clubs
		9 Three of Clubs
		10 Four of Clubs
	*/

	// 'hand' and 'remainingCards' are two new references
	// which points at two subsections of the 'cards' slice

	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}
