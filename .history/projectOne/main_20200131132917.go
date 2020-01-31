package main

import "fmt"

func main() {

	// ceate the deck of cards
	cardDeck := createDeck()
	//iterate
	cardDeck.print()

	// or card := newCard("hello")
	fmt.Println(card)
}

/*
Example of simple function use
func newCard(value string) string {
	return value
}
*/
