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

	fmt.Printf("%+v\n", mukul)

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@example.com",
			zipCode: 54321,
		},
	}

	fmt.Printf("%+v\n", alex)

}
