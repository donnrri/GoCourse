package main

import "fmt"

func main() {
	var card string = newCard("hello")

	cardDeck := deck{"one", "two"}
	//returns a new deck not mofiy current deck
	cardDeck = append(cardDeck, "three")
	//iterate
	for i, card := range cardDeck {
		fmt.Println(i, card)
	}

	// or card := newCard("hello")
	fmt.Println(card)
}

func newCard(value string) string {
	return value
}
