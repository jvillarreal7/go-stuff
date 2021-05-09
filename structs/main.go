package main

import "fmt"

/*
	Value types: int, float, string, bool, structs - use pointers to change these within a func
	Reference types: slices, maps, channels, pointers, functions - don't worry about pointers
*/

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 34532,
		},
	}
	jim.updateName("Jimothy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
