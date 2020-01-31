package main

// multiple imports of libs
import (
	"fmt"
	"ioutil"
	"strings"
)

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

func (d deck) saveToFile(filename string) error {

	// deck ([]string) -> string (toString)
	str := toString(d)

	// convert your slice to bytes (byte slice)
	// []byte(value) - type conversion
	bites := []byte(str)

	//use ioutils package to write file
	// retuen eror if it occurs
	return ioutil.WriteFile(filename, bites, 0666)

}

// helper function
func (d deck) toString() string {

	// type conversion
	converted := []string(d)

	//generate a string with comma separating elements of []string
	// use a packge lib 'strings'
	// Join(a []string, sep string) string (sep = separator)
	return strings.Join(converted, ",")

}
