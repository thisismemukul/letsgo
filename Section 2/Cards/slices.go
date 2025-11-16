package main

import "fmt"

func slice() {
	cards := []string{"1", newCard()}
	cards = append(cards, "2")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
