package main

import (
	"os"
	"testing"
)

func TestNewDeck2(tester *testing.T) {
	cards := newDeck()

	expected_length := 52
	if len(cards) != expected_length {
		tester.Errorf("Expected deck length of %v, but got %v", expected_length, len(cards))
	}

	expected_first_card := "Ace of Spades"
	if cards[0] != expected_first_card {
		tester.Errorf("Expected first card of %v, but got %v", expected_first_card, cards[0])
	}

	expected_last_card := "King of Clubs"
	if cards[len(cards)-1] != expected_last_card {
		tester.Errorf("Expected last card of %v, but got %v", expected_last_card, cards[len(cards)-1])
	}
}

func TestSaveToAndLoadFromFile(tester *testing.T) {
	// func (cards deck) saveToFile(fileName string) error
	// func newDeckFromFile(fileName string) deck
	// remove leftovers
	file_name_template := "_new_deck_save_load_testing.deck"
	os.Remove(file_name_template)
	// save to testing file
	cards := newDeck()
	err := cards.saveToFile(file_name_template)
	if err != nil {
		tester.Errorf("Error saving deck to file: %v", err)
	}
	// loading saved testing deck
	loaded_deck := newDeckFromFile(file_name_template)
	expected_length := 52
	if len(loaded_deck) != expected_length {
		tester.Errorf("Expected loaded from disk deck length of %v, but got %v", expected_length, len(loaded_deck))
	}

	if !isSameDeck(cards, loaded_deck) {
		tester.Errorf("Expected the loaded deck to be the same as the saved deck, but it was different")
	}
	// remove the testing deck
	os.Remove(file_name_template)
}

func TestShuffle(tester *testing.T) {
	starting_deck := newDeck()
	another_deck := make(deck, len(starting_deck))
	copy(another_deck, starting_deck)

	starting_deck.shuffle()

	if isSameDeck(starting_deck, another_deck) {
		tester.Errorf("Expected the deck to be shuffled, but it was the same as the original deck")
	}
}

func TestDeal(tester *testing.T) {
	cards := newDeck()
	handSize := 5
	hand, remainingDeck := deal(cards, handSize)

	if len(hand) != handSize {
		tester.Errorf("Expected hand size of %v, but got %v", handSize, len(hand))
	}

	if len(remainingDeck) != len(cards)-handSize {
		tester.Errorf("Expected remaining deck size of %v, but got %v", len(cards)-handSize, len(remainingDeck))
	}
}

func isSameDeck(deck1, deck2 deck) bool {
	if len(deck1) != len(deck2) {
		return false
	}
	for i := range deck1 {
		if deck1[i] != deck2[i] {
			return false
		}
	}
	return true
}
