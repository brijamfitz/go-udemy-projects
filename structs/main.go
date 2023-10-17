package main

import (
	"fmt"
)

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
	phil := person{
		firstName: "Phil",
		lastName:  "Mickelson",
		contactInfo: contactInfo{
			email:   "phil@mickelson.com",
			zipCode: 11111,
		},
	}

	phil.updateName("Ernie")
	phil.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
