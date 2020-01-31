package main

import "fmt"

func main() {

	// ceate the deck of cards
	cardDeck := createDeck()

	//iterate through and print out whole deck
	//cardDeck.print()
	// use the deal function
	first, remaining := deal(cardDeck, 5)
	// returns two decks (hands) of cards
	first.print()
	fmt.Println("and")
	remaining.print()

	cardDeck.saveToFile("projectOne/saved_files/my_cards")

}

/*
Example of simple function use
func newCard(value string) string {
	return value
}
*/
