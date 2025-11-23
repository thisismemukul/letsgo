package main

import "fmt"

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
	mukul := person{
		firstName: "Mukul",
		lastName:  "Saini",
		contact: contactInfo{
			email:   "mukul@example.com",
			zipCode: 12345,
		},
	}
	mukul.print()

	fmt.Printf("%+v\n", mukul)

	mukul.updateName("Mukul Saini")
	mukul.print()

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@example.com",
			zipCode: 54321,
		},
	}
	alex.print()

	fmt.Printf("%+v\n", alex)

	// alexPointer := &alex
	// alexPointer.updateName("Alex Anderson")
	alex.updateName("Alex Anderson")
	alex.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
