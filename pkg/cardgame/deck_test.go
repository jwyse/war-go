package cardgame

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	options := DeckCreationOptions{AceHigh: false}
	deck := CreateDeck(options)

	if deck.Len() != 52 {
		t.Fatalf("expected 52 cards but got %d", deck.Len())
	}

	foundAceLow := false
	foundAceHigh := false
	for deck.Len() > 0 {
		card := deck.DealCard()
		if card.rank == 1 {
			foundAceLow = true
		}
		if card.rank == 14 {
			foundAceHigh = true
		}
	}
	if foundAceHigh {
		t.Fatalf("created deck with Ace low, but found a card with rank 14 (Ace high)")
	}
	if !foundAceLow {
		t.Fatalf("created deck with Ace low, but didn't find a card with rank 1 (Ace low)")
	}

	options = DeckCreationOptions{AceHigh: true}
	deck = CreateDeck(options)

	if deck.Len() != 52 {
		t.Fatalf("expected 52 cards but got %d", deck.Len())
	}

	foundAceLow = false
	foundAceHigh = false
	for deck.Len() > 0 {
		card := deck.DealCard()
		if card.rank == 1 {
			foundAceLow = true
		}
		if card.rank == 14 {
			foundAceHigh = true
		}
	}
	if foundAceLow {
		t.Fatalf("created deck with Ace high, but found a card with rank 1 (Ace low)")
	}
	if !foundAceHigh {
		t.Fatalf("created deck with Ace high, but didn't find a card with rank 14 (Ace high)")
	}
}
