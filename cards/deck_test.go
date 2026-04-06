package main

import (
	"os"
	"testing"
)

// whenever writing testing code with go, handle cleanup urself
// file naming convention "_test" followed by whatever name u want
func TestNewDec(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12 but got %v", len(d)) // a formated string
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card: Ace of Spades, but got: %v", d[0])
	}

	if d[len(d)-1] != "Three of Clubs" {
		t.Errorf("Expected last card: Three of Clubs, but got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // if present remove

	deck := newDeck()

	deck.saveToFile("_decktesting")
}
