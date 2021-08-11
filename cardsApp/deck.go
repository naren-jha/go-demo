package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"} // we can add rest of the values

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/*
newDeck() function returns a deck, but doesn't need a receiver
because when you call this function, you actually want a deck (you
	don't have one already)

if you don't need the index in for loop, declare it as _
*/

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
deck can be thought of as an object which has print method
available on it. it's just like a class delclaration in
languages like java. Go is not an object oriented language,
but for simplicity, we can think of it in terms OOL.

so deck is a type (class) which basically represents a list (slice)
of string. and this type has print method available on it.

In above function, (d deck) means this function can be called on
deck type variables. so (d deck) is the receiver type object.
*/

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
