package main

import (
	"math/rand"
	"time"
)

type Deck []Card

type DeckCreationOptions struct {
	// Whether to shuffle the deck
	Shuffle bool
	// Ace high (...JQKA) or low (A234...)
	AceHigh bool
}

// Create a deck of standard playing cards with the specified options
func CreateDeck(options DeckCreationOptions) Deck {
	rand.Seed(time.Now().UnixNano())
	deck := Deck{}
	suits := GetSuits()
	for _, s := range suits {
		for i := 1; i < 14; i++ {
			card := Card{rank: uint8(i), suit: suits[s]}
			if options.AceHigh && i == 1 {
				card.rank = 14
			}
			deck.AddCard(card)
		}
	}
	if options.Shuffle {
		deck.Shuffle()
	}
	return deck
}

// Returns the number of cards in the deck
func (d Deck) len() int { return len(d) }

func (d Deck) String() string {
	s := ""
	for _, card := range d {
		s += card.String()
		s += " "
	}
	return s
}

// Shuffle the deck
func (d Deck) Shuffle() {
	for i := range d {
		d.swapCards(i, rand.Intn(d.len()))
	}
}

// Add a card to the bottom of a face-down deck
func (d *Deck) AddCard(card Card) {
	*d = append(*d, card)
}

// Add cards to the bottom of a face-down deck
func (d *Deck) addCards(cards ...Card) {
	for _, card := range cards {
		d.AddCard(card)
	}
}

// Deal a card from the top of a face-down deck
func (d *Deck) DealCard() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// Swap two cards (by index) within the deck
func (d Deck) swapCards(a, b int) {
	d[a], d[b] = d[b], d[a]
}

// Deal the specified number of cards to the specified number of players. Returns the player hands as Decks.
func (d *Deck) DealHands(players int, cardsPerPlayer int) []Deck {
	hands := make([]Deck, players)

	dealToPlayer := 0
	for i := 0; i < cardsPerPlayer*players; i++ {
		if d.len() == 0 { // No more cards to deal
			break
		}

		hands[dealToPlayer].AddCard(d.DealCard())

		dealToPlayer++
		if dealToPlayer >= len(hands) {
			dealToPlayer = 0
		}
	}

	return hands
}
