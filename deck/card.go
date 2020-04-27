//go:generate stringer -type=Rank,Suit

package deck

import (
	"fmt"
	"sort"
)

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)
const (
	minRank = Ace
	maxRank = King
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Card struct {
	Suit Suit
	Rank Rank
}

func New(options ...func([]Card) []Card) []Card {
	var deck []Card
	for _, suit := range suits {
		for i := minRank; i <= maxRank; i++ {
			deck = append(deck, Card{suit, i})
		}
	}
	for _, opt := range options {
		deck = opt(deck)
	}
	return deck
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cardValue(cards[i]) < cardValue(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func cardValue(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
