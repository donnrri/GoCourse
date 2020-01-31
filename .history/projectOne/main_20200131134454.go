package main

import "fmt"

func main() {

	// ceate the deck of cards
	cardDeck := createDeck()

	//iterate through and print out whole deck
	//cardDeck.print()

	first, remaining := deal(cardDeck, 5)

	first.print()
	fmt.Println("and")
	remaining.print()

}

/*
Example of simple function use
func newCard(value string) string {
	return value
}
*/
