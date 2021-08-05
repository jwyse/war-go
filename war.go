package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("\n    W A R !\n\n")
	options := DeckCreationOptions{
		Shuffle: true,
		AceHigh: false}
	deck := CreateDeck(options)
	playerHands := dealHands(2, deck)

	// The player who takes the hand leads the next hand. Start with Player 1
	leadPlayer := 0
	winner := -1
	hand := 0
	pot := Deck{}
	for {
		if len(playerHands[0]) == 0 {
			winner = 1
		} else if len(playerHands[1]) == 0 {
			winner = 0
		}

		if winner >= 0 {
			fmt.Printf("Player %d wins, after %d hands!", winner+1, hand)
			break
		}

		hand++
		p1card := playerHands[0].DealCard()
		p2card := playerHands[1].DealCard()
		if leadPlayer == 0 {
			fmt.Printf("Player 1: %s ; ", p1card)
			fmt.Printf("Player 2: %s ; ", p2card)
			pot.addCards(p1card, p2card)
		} else {
			fmt.Printf("Player 2: %s ; ", p2card)
			fmt.Printf("Player 1: %s ; ", p1card)
			pot.addCards(p2card, p1card)
		}

		handWinner := -1
		switch compare(p1card, p2card) {
		case -1:
			{
				// p1 beats p2
				handWinner = 0
			}
		case 1:
			{
				// p2 beats p1
				handWinner = 1
			}
		default:
			{
				// WAR!
				fmt.Printf("WAR !!\n\n")
				// time.Sleep(time.Duration(time.Millisecond) * 500)
				// pot.addCards(p1card, p2card)
				continue
			}
		}

		fmt.Printf("-> Player %d wins the hand!\n\n", handWinner+1)
		for {
			if pot.len() == 0 {
				break
			}
			playerHands[handWinner].AddCard(pot.DealCard())
		}
		// if leadPlayer == 1 {
		// 	// P2 led this hand, so add their card first
		// 	playerHands[handWinner].addCards(p2card, p1card)
		// } else {
		// 	playerHands[handWinner].addCards(p1card, p2card)
		// }

		// player who wins the hand leads the next hand
		leadPlayer = handWinner
		showHands(playerHands[0], playerHands[1])
	}
}

func showHands(p1, p2 Deck) {
	fmt.Printf("Player 1: (%d) %s\n", p1.len(), p1)
	fmt.Printf("Player 2: (%d) %s\n\n", p2.len(), p2)
	if p1.len()+p2.len() != 52 {
		panic("Should be 52 cards in play!")
	}
}

func dealHands(players int, deck Deck) []Deck {
	hands := make([]Deck, players)

	dealToPlayer := 0
	for {
		if deck.len() == 0 { // No more cards to deal
			break
		}

		hands[dealToPlayer].AddCard(deck.DealCard())

		dealToPlayer++
		if dealToPlayer >= len(hands) {
			dealToPlayer = 0
		}

		// fmt.Printf("Deck: %s\n", deck)
		// fmt.Printf("P1: %s (%d)\n", hands[0], hands[0].len())
		// fmt.Printf("P2: %s (%d)\n", hands[1], hands[1].len())
		// fmt.Println()
	}

	return hands
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

func compare(c1, c2 Card) int8 {
	if c1.rank == c2.rank {
		return 0
	} else if c1.rank > c2.rank {
		return -1
	} else {
		return 1
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
			card := Card{rank: uint8(i), suit: suits[s]}
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
	s := ""
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

func (d *Deck) addCards(cards ...Card) {
	for _, card := range cards {
		d.AddCard(card)
	}
}

func (d *Deck) DealCard() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

type Card struct {
	rank uint8
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
	//  1=Ace   => A  (Ace low)
	// 11=Jack  => J
	// 12=Queen => Q
	// 13=King  => K
	// 14=Ace   => A  (Ace high)
	if i >= 2 && i <= 10 {
		return fmt.Sprint(i)
	}
	rank := Rank(i)
	return string(rank[0])
}

func (c Card) Name() string {
	value := Rank(uint8(c.rank))
	return fmt.Sprintf("%v of %v", value, c.suit)
}

func (c Card) String() string {
	value := RankAbbrev(uint8(c.rank))
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
	if 40 == len(suit) {
		panic("nothing")
	}
	return fmt.Sprintf("%v%v", value, c.suit)
}
