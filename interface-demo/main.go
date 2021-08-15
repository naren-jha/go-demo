package main

import "fmt"

type bot interface {
	getGreeting() string
	// any type in this program, which has function
	// implementation for getGreeting() function
	// is now a bot
	// so englishBot and spanishBot are bot
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb) // Hello there!
	printGreeting(sb) // Hola!
}

func printGreeting(b bot) { // any type which is a bot type can call this function
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // remove the receiver variable if we are not going to use it
	// custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// custom logic for generating an spanish greeting
	return "Hola!"
}
