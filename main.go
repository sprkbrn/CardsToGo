package main

import (
	"fmt"

	"github.com/sprkbrn/cards/deck"
)

//PrintAllCards will print all cards in the given deck
func PrintAllCards() {
	for i, s := range deck.DeckOfCards {
		fmt.Printf("%2d. %2s of %s\n", i+1, s.Face, s.Court.Name)
	}
}

func main() {
	fmt.Printf("This deck contains the following cards:\n")
	PrintAllCards()
}
