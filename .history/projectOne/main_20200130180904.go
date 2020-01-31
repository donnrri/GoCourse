package main

import "fmt"

func main() {
	var card string = newCard("hello")
	// or card := "Something"
	fmt.Println(card)
}

func newCard(value string) string {
	return value
}
