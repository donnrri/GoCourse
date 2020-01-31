package main

import "fmt"

func main() {
	var card string = newCard("hello")

	deck := []string{newCard("one"), newCard("two")}
	deck = append(deck, "three")
	// or card := newCard("hello")
	fmt.Println(card)
}

func newCard(value string) string {
	return value
}
