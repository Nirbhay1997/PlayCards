package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Error("No of cards not 16,Got ", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Error("First card shoud be Ace of Spade")
	}

}

func TestSaveToDeckAndNewDeckFormFile(t *testing.T) {
	os.Remove("_decktest")
	d := newDeck()
	d.shuffle()
	d.saveToFile("_decktest")

	loaded_deck := newDeckFromFile("_decktest")

	if len(loaded_deck) != 16 {
		t.Error("No of cards not 16,Got ", len(d))
	}
	os.Remove("_decktest")

}
