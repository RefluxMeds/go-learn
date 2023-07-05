package main

import (
	"os"
	"testing"
)

// t.Errorf("Expected deck length of 16, but got", len(d))
// t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("The lenght of a new deck is less than 52!\tGot: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Want: Ace of Spades\tGot: %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Want: King of Clubs\tGot: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("The lenght of a new deck is less than 52!\tGot: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
