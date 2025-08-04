package main

import "github.com/dhananjayshk/card-deck/deck"

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// cards.saveToFile(("my_cards"))
	// cards, _ := newDeckFromFile("my_cards")/

	cards := deck.NewDeck()

	cards.Shuffle()
	cards.Print()
}
