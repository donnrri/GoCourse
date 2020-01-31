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
 1. type deck has access to the print function
 2. iterate throug the deck
*/

func (d deck) print() { //1
	for i, card := range d {
		2 //
		fmt.Println(i, card)
	}
}
