package main

import "fmt"

func main() {
	var card string = newCard("hello")

	cardDeck := deck[]string{newCard("one"), newCard("two")}
	//returns a new deck not mofiy current deck
	cardDeck = append(deck, "three")
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
