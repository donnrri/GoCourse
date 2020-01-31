package main

import "fmt"

func main() {
	var card string = newCard("hello")

	deck := []string{newCard("one"), newCard("two")}
	//returns a new deck not mofiy current deck
	deck = append(deck, "three")
	//iterate
	for i, card := range deck {
		fmt.Println(i, card)
	}

	// or card := newCard("hello")
	fmt.Println(card)
}

func newCard(value string) string {
	return value
}
