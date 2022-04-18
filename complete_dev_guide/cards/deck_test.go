package main

import (
	"os"
	"testing"
)

// tests has capital letters
// t *testing.T
func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected deck lentgh of 16 but found %v", len(deck))
	}

}

func TestSaveToDeck(t *testing.T) {
	testingPath := "./_decktesting"
	os.Remove(testingPath)
	deck := newDeck()
	deck.saveToFile(testingPath)
	newDeck := newDeckFromFile(testingPath)

	if len(newDeck) != 16 {
		t.Errorf("Expected deck lentgh of 16 but found %v", len(deck))
	}
	os.Remove(testingPath)
}
