package main

import "fmt"

func main() {
	var card string = newCard()
	// or card := "Something"
	fmt.Println(card)
}

func newCard() string {
	return "King of Hearts"
}
