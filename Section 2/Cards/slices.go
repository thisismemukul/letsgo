package main

func slice() {
	cards := deck{"1", newCard()}
	cards = append(cards, "2")

	cards.print()
}
