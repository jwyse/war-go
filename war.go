package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("\n    W A R !\n\n")
	aceHigh := false
	options := DeckCreationOptions{
		Shuffle: true,
		AceHigh: aceHigh,
	}
	deck := CreateDeck(options) // createDeck(aceHigh)
	// TODO: deal hands to players
	// decks := [2]Deck{createDeck(aceHigh), createDeck(aceHigh)}
	// for _, deck := range decks {
	// 	deck.Display()
	// }

	fmt.Println(deck)
	fmt.Printf("Deck has %d cards\n", deck.len())
	var dealtCard Card
	// Deal 3 cards
	for i := 0; i < 3; i++ {
		dealtCard = deck.DealCard()
		fmt.Printf("Dealt card %v\n", dealtCard)
		deck.AddCard(dealtCard)
	}
	fmt.Printf("Deck has %d cards\n", deck.len())
	fmt.Println(deck)
}

type Suit uint8

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

func (s Suit) String() string {
	switch s {
	case Hearts:
		{
			return "♥"
		}
	case Diamonds:
		{
			return "♦"
		}
	case Clubs:
		{
			return "♣"
		}
	case Spades:
		{
			return "♠"
		}
	default:
		panic("Unable to convert suit to string")
	}
}

type Deck []Card

type DeckCreationOptions struct {
	Shuffle bool
	AceHigh bool
}

func CreateDeck(options DeckCreationOptions) Deck {

	// Creates a deck of standard playing cards

	deck := Deck{}
	suits := []Suit{Spades, Hearts, Clubs, Diamonds}
	for _, s := range suits {
		for i := 1; i < 14; i++ {
			card := Card{rank: i, suit: suits[s]}
			if options.AceHigh && i == 1 {
				card.rank = 14
			}
			deck.AddCard(card)
		}
	}
	if options.Shuffle {
		deck.Shuffle(9)
	}
	return deck
}

func (d Deck) len() int { return len(d) }

func (d Deck) String() string {
	var s string
	for _, card := range d {
		s += card.String()
		s += " "
	}
	return s
}

func (d Deck) Shuffle(iterations int) {
	for i := 0; i < iterations; i++ {
		for range d {
			d.SwapCards(i, rand.Intn(d.len()))
		}
	}
}

func (d Deck) SwapCards(a, b int) {
	d[a], d[b] = d[b], d[a]
}

func (d *Deck) AddCard(card Card) {
	*d = append(*d, card)
}

func (d *Deck) DealCard() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

type Card struct {
	rank int
	suit Suit
}

func Rank(i uint8) string {

	// convert 1/11-13 to A/JQKA

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

func RankAbbrev(i uint8) string {
	rank := Rank(i)
	return string(rank[0])
}

func (c Card) Name() string {
	value := Rank(uint8(c.rank))
	return fmt.Sprintf("%v of %v", value, c.suit)
}

func (c Card) String() string {
	value := RankAbbrev(uint8(c.rank))
	// value := fmt.Sprint(c.rank)
	// switch c.rank {
	// case 1:
	// 	{
	// 		value = "A"
	// 	}
	// case 11:
	// 	{
	// 		value = "J"
	// 	}
	// case 12:
	// 	{
	// 		value = "Q"
	// 	}
	// case 13:
	// 	{
	// 		value = "K"
	// 	}
	// case 14:
	// 	{
	// 		value = "A"
	// 	}
	// }
	suit := "♦♣♠♥"
	switch c.suit {
	case Hearts:
		{
			suit = "♥"
		}
	case Diamonds:
		{
			suit = "♦"
		}
	case Clubs:
		{
			suit = "♣"
		}
	case Spades:
		{
			suit = "♠"
		}
	}
	return fmt.Sprintf("%v%v", value, suit)
}
