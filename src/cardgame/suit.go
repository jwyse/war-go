package cardgame

import "fmt"

type Suit uint8

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

func GetSuits() []Suit {
	return []Suit{Spades, Hearts, Clubs, Diamonds}
}

func (s Suit) String() string {
	return s.Symbol()
}

func (s Suit) Name() string {
	switch s {
	case Spades:
		{
			return "Spades"
		}
	case Hearts:
		{
			return "Hearts"
		}
	case Clubs:
		{
			return "Clubs"
		}
	case Diamonds:
		{
			return "Diamonds"
		}
	default:
		panic(fmt.Sprintf("Unable to convert suit %d to string", s))
	}
}

func (s Suit) Symbol() string {
	switch s {
	case Spades:
		{
			return "♠"
		}
	case Hearts:
		{
			return "♥"
		}
	case Clubs:
		{
			return "♣"
		}
	case Diamonds:
		{
			return "♦"
		}
	default:
		panic(fmt.Sprintf("Unable to convert suit %d to string", s))
	}
}
