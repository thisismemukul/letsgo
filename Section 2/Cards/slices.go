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
}
