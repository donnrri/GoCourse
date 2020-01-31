package main

import "fmt"
/*
Create a new type of deck which has a slice of strings
Functions to manipulate the deck are attached to it
It is roughly equivilent to a class
*/

type deck []string

func (d deck) print{
	for i, card := range d {
		fmt.Println(i, card)
	}
}
