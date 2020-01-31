package main

import "fmt"

func main() {
	var card string = newCard("hello")

	deck := newDeck()
	// or card := newCard("hello")
	fmt.Println(card)
}

func newCard(value string) string {
	return value
}

func newDeck() {
	return []string{newCard("one"), newCard("two")}
}
