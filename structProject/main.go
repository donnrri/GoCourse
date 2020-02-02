package main

import "fmt"

// person has a nested struct type
type person struct {
	name    string
	age     int
	contact contactDetails
}

type contactDetails struct {
	email string
	zip   string
}

func main() {

	// tony := person{"Tony", 90} uses order
	//or using key value approach
	tony := person{name: "Tony", age: 90}
	fmt.Println(tony.name)

	// or
	// here all values e.g name  get zero values
	//name = '', int = 0 , bool = false
	var alex person
	contact := contactDetails{
		email: "me@gmail.com",
		zip:   "eh87fg",
	}
	alex.contact = contact

	alexPointer := &alex
	alexPointer.updateName("newName")

}

/*
 func takes a struct as a recievr ,
 the name is not updated in the person passed in

func (p person) updateName(newName string) {
	p.name = newName
}
*/

/*
 funcs can take structs as pointers
 in this case th value is updated
 Notice the very different syntax
*/
func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).name = newName
}

/*
Pointer operations:

alexPointer = &alex
'&' is an operator this says give me access to the memory
 slot of this variable. Its like a reference to the var
 This is assigned to alexPointer
(*pointerToPerson) = give me access to the value and
I can now modify it.


*/
