package main

import (
	"fmt"

	"github.com/jwyse/war/pkg/cardgame"
)

func main() {
	fmt.Printf("\n    W A R !\n\n")
	options := cardgame.DeckCreationOptions{
		Shuffle: true,
		AceHigh: false}
	deck := cardgame.CreateDeck(options)
	fmt.Printf("Initial deck: %s\n\n", deck)
	playerHands := deck.DealHands(2, 26)
	validateDecks(deck, playerHands[0], playerHands[1])

	// The player who takes the hand leads the next hand. Start with Player 1
	leadPlayer := 0
	winner := -1
	hand := 0
	pot := cardgame.Deck{}
	for {
		if len(playerHands[0]) == 0 {
			winner = 1
		} else if len(playerHands[1]) == 0 {
			winner = 0
		}

		if winner >= 0 {
			fmt.Printf("Player %d wins, after %d hands!\n", winner+1, hand)
			break
		}

		hand++
		p1card := playerHands[0].DealCard()
		p2card := playerHands[1].DealCard()
		if leadPlayer == 0 {
			fmt.Printf("Player 1: %s ; Player 2: %s ; ", p1card, p2card)
			pot.AddCards(p1card, p2card)
		} else {
			fmt.Printf("Player 2: %s ; Player 1: %s ; ", p2card, p1card)
			pot.AddCards(p2card, p1card)
		}

		handWinner := -1
		switch cardgame.Compare(p1card, p2card) {
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
				continue
			}
		}

		fmt.Printf("-> Player %d wins the hand!\n\n", handWinner+1)
		for {
			if pot.Len() == 0 {
				break
			}
			playerHands[handWinner].AddCard(pot.DealCard())
		}

		// player who wins the hand leads the next hand
		leadPlayer = handWinner
		showHands(deck, playerHands[0], playerHands[1])
	}
}

func showHands(deck, p1, p2 cardgame.Deck) {
	fmt.Printf("Player 1: (%d) %s\n", p1.Len(), p1)
	fmt.Printf("Player 2: (%d) %s\n\n", p2.Len(), p2)
	validateDecks(deck, p1, p2)
}

func validateDecks(deck, p1, p2 cardgame.Deck) {
	if deck.Len() != 0 {
		panic("Deck should be empty after dealing to players")
	}
	if p1.Len()+p2.Len() != 52 {
		panic("There should be exactly 52 cards in play")
	}
}
