package main

import "testing"
import "os"
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck to be 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[16])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_deckstring")

	loadedDeck := newDeckFromFile("_deckstring")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of deck to be 16, but got %v", len(loadedDeck))
	}
	os.Remove("_deckstring")
}