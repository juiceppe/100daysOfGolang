package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Lenght expected 16, but instead: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace Of Spades but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card Four Of Clubs but got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTest")
	d := newDeck()
	d.saveToFile("_deckTest")

	loadedDeck := newDeckFromFile("_deckTest")

	if len(loadedDeck) != 16 {
		t.Errorf("Lenght expected 16, but instead: %v", len(loadedDeck))
	}
	os.Remove("_deckTest")
}
