package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()

	if len(d) != 16 {
		t.Errorf("Expected d of length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fileName := fmt.Sprintf("%v/_decktesting%v", "Decks", ".txt")

	os.Remove(fileName)

	deck := createDeck()

	deck.saveFile("_decktesting")

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected d of length 16, but got %v", len(deck))
	}

	os.Remove("_decktesting" + ".txt")
}
