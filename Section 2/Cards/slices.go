package main

import "fmt"

func slice() {
	cards := newDeck()
	// cards = append(cards, "2")
	hand, remaining := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining:")
	remaining.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
	newDeck := newDeckFromFile("my_cards.txt")
	newDeck.print()
	newDeck.shuffle()
	newDeck.print()
}
