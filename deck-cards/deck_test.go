package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := createDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}
}
