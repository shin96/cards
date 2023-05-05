package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length to be 16, got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected card to be 'Ace of Diamonds'but found %v", d[0])
	}

}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")

    loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected deck length to be %v, got %v", len(deck), len(loadedDeck))
	}

	os.Remove("_deckTesting")

}
