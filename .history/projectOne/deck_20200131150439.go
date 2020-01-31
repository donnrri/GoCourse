package main

import "fmt"

/*
Create a new type of deck which has a slice of strings
Functions to manipulate the deck are attached to it
It is roughly equivilent to a class
*/

type deck []string

/*
 this is a special kind of function called a receiver.
 1. type deck has access to the print function - note convention to
 	have the var as a single letter , here 'd'
 2. iterate through the deck
*/
//1
func (d deck) print() {
	//2
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func createDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// if not using variable like i and j
	// replace with underscores
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of  "+suit)
		}
	}

	return cards

}

/*
	(d deck, handSize int) two argumants
	(deck, deck) - indicates two values are returned from unc
	return d[:handSize], d[handSize:] returning te list
	split on handsize
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func save(d deck) {

	// convert your slice to bytes (byte slice)
	// []byte (value)

}
