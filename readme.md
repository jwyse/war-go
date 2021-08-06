# War (card game) implemented in Go

War is a card "game" that doesn't involve any decisions or strategy from the player - the winner is determined by the shuffle/deal and which player goes first.  I'm implementing it as an exercise to become more familiar with the Go language.

Slight modification to the Bicycle rules below: when "war" occurs, each player simply turns over one additional card.

Sample output:
```
./war-go.exe

    W A R !

Initial deck: 10♦ 3♥ K♠ 3♠ 10♣ 9♥ 7♥ 6♣ K♥ 4♣ 8♣ 3♣ Q♦ Q♠ 4♦ 4♠ 2♦ A♠ 7♠ A♣ 2♣ 6♠ 9♦ 9♠ 7♦ J♣ Q♣ Q♥ A♥ 8♦ K♦ 9♣ 8♥ J♠ 7♣ 4♥ 8♠ 6♦ 2♠ 3♦ 2♥ K♣ 5♦ 5♥ A♦ J♥ 5♠ 6♥ 10♠ J♦ 10♥ 5♣

Player 1: 10♦ ; Player 2: 3♥ ; -> Player 1 wins the hand!

Player 1: (27) K♠ 10♣ 7♥ K♥ 8♣ Q♦ 4♦ 2♦ 7♠ 2♣ 9♦ 7♦ Q♣ A♥ K♦ 8♥ 7♣ 8♠ 2♠ 2♥ 5♦ A♦ 5♠ 10♠ 10♥ 10♦ 3♥
Player 2: (25) 3♠ 9♥ 6♣ 4♣ 3♣ Q♠ 4♠ A♠ A♣ 6♠ 9♠ J♣ Q♥ 8♦ 9♣ J♠ 4♥ 6♦ 3♦ K♣ 5♥ J♥ 6♥ J♦ 5♣

Player 1: K♠ ; Player 2: 3♠ ; -> Player 1 wins the hand!

Player 1: (28) 10♣ 7♥ K♥ 8♣ Q♦ 4♦ 2♦ 7♠ 2♣ 9♦ 7♦ Q♣ A♥ K♦ 8♥ 7♣ 8♠ 2♠ 2♥ 5♦ A♦ 5♠ 10♠ 10♥ 10♦ 3♥ K♠ 3♠
Player 2: (24) 9♥ 6♣ 4♣ 3♣ Q♠ 4♠ A♠ A♣ 6♠ 9♠ J♣ Q♥ 8♦ 9♣ J♠ 4♥ 6♦ 3♦ K♣ 5♥ J♥ 6♥ J♦ 5♣

...

Player 1: (32) Q♦ 4♦ 2♦ 7♠ 2♣ 9♦ 7♦ Q♣ A♥ K♦ 8♥ 7♣ 8♠ 2♠ 2♥ 5♦ A♦ 5♠ 10♠ 10♥ 10♦ 3♥ K♠ 3♠ 10♣ 9♥ 7♥ 6♣ K♥ 4♣ 8♣ 3♣
Player 2: (20) Q♠ 4♠ A♠ A♣ 6♠ 9♠ J♣ Q♥ 8♦ 9♣ J♠ 4♥ 6♦ 3♦ K♣ 5♥ J♥ 6♥ J♦ 5♣

Player 1: Q♦ ; Player 2: Q♠ ; WAR !!

Player 1: 4♦ ; Player 2: 4♠ ; WAR !!

Player 1: 2♦ ; Player 2: A♠ ; -> Player 1 wins the hand!

Player 1: (35) 7♠ 2♣ 9♦ 7♦ Q♣ A♥ K♦ 8♥ 7♣ 8♠ 2♠ 2♥ 5♦ A♦ 5♠ 10♠ 10♥ 10♦ 3♥ K♠ 3♠ 10♣ 9♥ 7♥ 6♣ K♥ 4♣ 8♣ 3♣ Q♦ Q♠ 4♦ 4♠ 2♦ A♠
Player 2: (17) A♣ 6♠ 9♠ J♣ Q♥ 8♦ 9♣ J♠ 4♥ 6♦ 3♦ K♣ 5♥ J♥ 6♥ J♦ 5♣

...

Player 1: (48) 3♦ 8♥ K♠ 9♣ K♣ 8♣ J♥ 10♦ Q♥ 4♥ 7♠ 2♥ Q♦ A♦ 2♣ 6♣ 2♦ A♣ 3♣ 9♥ 5♥ 3♥ 5♠ 5♣ 8♦ 7♦ 4♣ 10♠ K♥ J♣ 6♦ A♠ J♦ 10♣ Q♠ 7♣ 5♦ 6♥ 4♠ 4♦ 6♠ K♦ Q♣ 9♠ 9♦ 7♥ 10♥ A♥
Player 2: (4) 3♠ 8♠ J♠ 2♠

Player 2: 3♠ ; Player 1: 3♦ ; WAR !!

Player 2: 8♠ ; Player 1: 8♥ ; WAR !!

Player 2: J♠ ; Player 1: K♠ ; -> Player 1 wins the hand!

Player 1: (51) 9♣ K♣ 8♣ J♥ 10♦ Q♥ 4♥ 7♠ 2♥ Q♦ A♦ 2♣ 6♣ 2♦ A♣ 3♣ 9♥ 5♥ 3♥ 5♠ 5♣ 8♦ 7♦ 4♣ 10♠ K♥ J♣ 6♦ A♠ J♦ 10♣ Q♠ 7♣ 5♦ 6♥ 4♠ 4♦ 6♠ K♦ Q♣ 9♠ 9♦ 7♥ 10♥ A♥ 3♠ 3♦ 8♠ 8♥ J♠ K♠
Player 2: (1) 2♠

Player 1: 9♣ ; Player 2: 2♠ ; -> Player 1 wins the hand!

Player 1: (52) K♣ 8♣ J♥ 10♦ Q♥ 4♥ 7♠ 2♥ Q♦ A♦ 2♣ 6♣ 2♦ A♣ 3♣ 9♥ 5♥ 3♥ 5♠ 5♣ 8♦ 7♦ 4♣ 10♠ K♥ J♣ 6♦ A♠ J♦ 10♣ Q♠ 7♣ 5♦ 6♥ 4♠ 4♦ 6♠ K♦ Q♣ 9♠ 9♦ 7♥ 10♥ A♥ 3♠ 3♦ 8♠ 8♥ J♠ K♠ 9♣ 2♠
Player 2: (0)

Player 1 wins, after 176 hands!
```

[Rules, according to Bicycle:](https://bicyclecards.com/how-to-play/war/)

The goal is to be the first player to win all 52 cards

### The Deal

The deck is divided evenly, with each player receiving 26 cards, dealt one at a time, face down. Anyone may deal first. Each player places their stack of cards face down, in front of them.

### The Play

Each player turns up a card at the same time and the player with the higher card takes both cards and puts them, face down, on the bottom of his stack.

If the cards are the same rank, it is War. Each player turns up one card face down and one card face up. The player with the higher cards takes both piles (six cards). If the turned-up cards are again the same rank, each player places another card face down and turns another card face up. The player with the higher card takes all 10 cards, and so on.

### How to Keep Score

The game ends when one player has won all the cards.