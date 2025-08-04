package deck_test

import (
	"os"
	"testing"

	"github.com/dhananjayshk/card-deck/deck"
)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := deck.NewDeck()
	deck.SaveToFile("_decktesting")
	loadedDeck, err := deck.NewDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

	if err != nil {
		t.Errorf("Error loading deck from file: %v", err)
	}

}
