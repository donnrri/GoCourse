package main

// multiple imports of libs
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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
	str := d.toString()

	// convert your slice to bytes (byte slice)
	// []byte(value) - type conversion
	bites := []byte(str)

	//use ioutils package to write file
	// return eror if it occurs
	return ioutil.WriteFile(filename, bites, 0644)

}

func (d deck) readFile(filename string) (deck, error) {
	byteSlice, err := ioutil.ReadFile(filename)
	//catch err
	if err != nil { // not sure this is needed
		os.Exit(1) // quit program
	}

	s := strings.Split(string(byteSlice), ",")
	return deck(s), err // again return two values

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

func (d deck) shuffle() deck {

	source := rand.NewSource()
	r := rand.New(source)
	for i := range d {
		newPos := rand.Intn(len(d) - 1) // randiom number generator
		d[i], d[newPos] = d[newPos], d[i]
	}

	return d
}
