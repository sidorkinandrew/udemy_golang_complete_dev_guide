package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("saved_deck.dck")

	old_cards := newDeckFromFile("saved_deck.dck")
	old_cards.print()

	// no_cards := newDeckFromFile("no-such-file.dck")
	// no_cards.print()

	cards.shuffle()
	cards.print()

}
