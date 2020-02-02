package main

import "fmt"

func main() {
	// declare map
	// here key of type string and
	// value of type string ([string]string)
	colors := map[string]string{
		"red":   "#ffff00",
		"green": "#4bf746",
	}

	/*
		Alternate declaration
		empty map
		colors := make(map[string]string)
		colors["black"] = "#000000"

	*/

	// acess a key / value use square bracket syntax
	fmt.Println(colors["red"])

	// delete a value
	delete(colors, "red")

	// iterate over a map
	for color, hex := range colors {
		fmt.Println(color, hex)
	}
}
