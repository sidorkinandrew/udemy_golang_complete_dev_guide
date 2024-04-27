package main

import "testing"

func TestNesDeck(tester *testing.T) {
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
