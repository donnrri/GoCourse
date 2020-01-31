package main

import "fmt"

func main() {
	var card string = newCard("hello")
	// or card := newCard("hello")
	fmt.Println(card)
}

func newCard(value string) string {
	return value
}

func newSlice() {
	cards := []string 

}
