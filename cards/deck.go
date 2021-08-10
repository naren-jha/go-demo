package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

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

so deck is type (class) which basically represents a list (slice)
of string. and this type has print method available on it.

In above function, (d deck) means this function can be called on
deck type variables.
*/
