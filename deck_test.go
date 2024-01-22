package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck len 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of ดอกจิก" {
		t.Errorf("Expected deck len Ace of ดอกจิก, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "K of โพธ์ดำ" {
		t.Errorf("Expected deck len Ace of ดอกจิก, but got %v", deck[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Load and SaveExpected")
	}
	os.Remove(fileName)
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	hand, remain := deal(deck, 5)
	if len(hand) != 5 {
		t.Errorf("Hand error")
	}
	if len(remain) != len(deck)-len(hand) {
		t.Errorf("remain error")
	}
}
