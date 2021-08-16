package cardgame

import "fmt"

type PlayingCard struct {
	rank uint8
	suit Suit
}

// Compare the ranks of two cards.
// Returns -1 if the first card is higher, 1 if the second card is higher, or 0 if the ranks are equal.
// Ace high/low was determined when the deck was created.
// Suit is ignored.
func Compare(c1, c2 PlayingCard) int8 {
	if c1.rank == c2.rank {
		return 0
	} else if c1.rank > c2.rank {
		return -1
	} else {
		return 1
	}
}

// Returns the rank name, i.e. "Ace", "Queen", "3"
func rank(i uint8) string {
	value := fmt.Sprint(i)
	switch i {
	case 1:
		{
			value = "Ace"
		}
	case 11:
		{
			value = "Jack"
		}
	case 12:
		{
			value = "Queen"
		}
	case 13:
		{
			value = "King"
		}
	case 14:
		{
			value = "Ace"
		}
	}
	return value
}

// Returns the rank abbreviation, i.e. "A" (Ace), "Q" (Queen), "3"
func rankAbbrev(i uint8) string {
	if i >= 2 && i <= 10 {
		return fmt.Sprint(i)
	}
	rank := rank(i)
	return string(rank[0])
}

// Returns the name of the card; i.e. "Ace of Spades", "2 of Clubs"
func (c PlayingCard) Name() string {
	value := rank(uint8(c.rank))
	return fmt.Sprintf("%v of %v", value, c.suit)
}

func (c PlayingCard) String() string {
	value := rankAbbrev(uint8(c.rank))
	return fmt.Sprintf("%v%v", value, c.suit)
}
