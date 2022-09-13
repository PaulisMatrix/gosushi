package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades,but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs,but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	// create a new deck and save to the file
	deck := NewDeck()
	deck.SaveToFile("_decktesting")

	//read from the file and assert
	loadedDeck := NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
