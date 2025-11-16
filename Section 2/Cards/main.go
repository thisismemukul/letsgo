package main

import "fmt"

var card string = "Ace of Spades"

// card := "Ace of Spades"

func main() {
	// card := "Ace of Spades"
	// card = "Five of diimonds"
	card = newCard()
	fmt.Println(card)
}

func newCard() string {
	return " new Card"
}
