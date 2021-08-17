package cardgame

import (
	"strconv"
	"testing"
)

func TestRankName(t *testing.T) {
	tests := make(map[uint8]string)
	var i uint8
	for i = 2; i < 11; i++ {
		tests[i] = strconv.Itoa(int(i))
	}

	tests[uint8(0)] = "0" // invalid
	tests[uint8(1)] = "Ace"
	tests[uint8(11)] = "Jack"
	tests[uint8(12)] = "Queen"
	tests[uint8(13)] = "King"
	tests[uint8(14)] = "Ace"
	tests[uint8(15)] = "15" // invalid

	for rank, want := range tests {
		have := rankName(rank)
		if have != want {
			t.Fatalf("expected card rank '%v' but got '%v'", want, rank)
		}
	}
}

func TestRankAbbrev(t *testing.T) {
	tests := make(map[uint8]string)
	var i uint8
	for i = 2; i < 11; i++ {
		tests[i] = strconv.Itoa(int(i))
	}

	tests[uint8(0)] = "0" // invalid
	tests[uint8(1)] = "A"
	tests[uint8(11)] = "J"
	tests[uint8(12)] = "Q"
	tests[uint8(13)] = "K"
	tests[uint8(14)] = "A"
	// invalid: tests[uint8(15)] = "15"

	for rank, want := range tests {
		have := rankAbbrev(rank)
		if have != want {
			t.Fatalf("%v: expected card rank abbrev '%v' but got '%v'", rank, want, have)
		}
	}
}
func TestCardName(t *testing.T) {
	tests := map[PlayingCard]string{
		// empty PlayingCard{} will panic -- {}: "",
		{rank: 1, suit: Clubs}:     "Ace of Clubs",
		{rank: 10, suit: Hearts}:   "10 of Hearts",
		{rank: 11, suit: Diamonds}: "Jack of Diamonds",
		{rank: 12, suit: Clubs}:    "Queen of Clubs",
		{rank: 13, suit: Spades}:   "King of Spades",
		{rank: 14, suit: Spades}:   "Ace of Spades",
	}

	for k, v := range tests {
		rankName(0)
		want := v
		have := k.Name()
		if have != want {
			t.Fatalf("name should be %v but was %v\n", want, have)
		}
	}
}
