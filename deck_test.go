package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expect deck length of 52, but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card Ace of Spade but got", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Error("Expect King of Clubs but got ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Error("Expected 52 card in deck, got ", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
