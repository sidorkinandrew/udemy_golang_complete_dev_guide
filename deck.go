package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of "deck"
// a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Valet", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(cards.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("[func newDeckFromFile(fileName string) deck]", fileName, "- error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bytes), ","))
}
