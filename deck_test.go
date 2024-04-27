package main

import "testing"

func TestNesDeck(tester *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		tester.Errorf("Expected deck length of 52, but got %v", len(cards))
	}
}
