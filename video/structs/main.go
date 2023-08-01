package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// ========== Alex ==========
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// ========== Alex Pointer ==========
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)

	// ========== Jim Person ==========

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// ========== Jim Pointer ==========
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.print()

	// ========== Pointer Shortcut ==========
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
