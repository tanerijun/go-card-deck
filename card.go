//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"sort"
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

type Rank uint8

const (
	_ Rank = iota
	Ace
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

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New returns a deck of cards.
//
// Optionally, you can pass functions to modify the generated deck (Sorting, shuffling, etc).
//
// The function format is:
//
//	func([]Card) []Card
//
// There're also some default helpers shipped with the API. For example: DefaultSort, Sort, etc.
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// DefaultSort sort cards according to their absolute rank.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, less(cards))
	return cards
}

func less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return AbsRank(cards[i]) < AbsRank(cards[j])
	}
}

// Sort sort cards with a custom Less function.
//
// Example:
// Sorting in reverse order
//
//	Sort(func(cards []Card) func(i, j int) bool {
//		return func(i, j int) bool { return AbsRank(cards[i]) > AbsRank((cards[j])) }
//	})
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Return a card's absolute rank.
//
// The rank of any Spade card will never be above Diamond.
func AbsRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
