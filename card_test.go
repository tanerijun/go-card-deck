package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Errorf("Expected New to return 52 cards, got %d", len(cards))
	}

	for i := 0; i < len(cards); i++ {
		if i <= 12 {
			if cards[i].Suit != Spade {
				t.Errorf("Expected Spade, got %s", cards[i].Suit)
			}
		} else if i <= 25 {
			if cards[i].Suit != Diamond {
				t.Errorf("Expected Diamond, got %s", cards[i].Suit)
			}
		} else if i <= 38 {
			if cards[i].Suit != Club {
				t.Errorf("Expected Club, got %s", cards[i].Suit)
			}
		} else {
			if cards[i].Suit != Heart {
				t.Errorf("Expected Heart, got %s", cards[i].Suit)
			}
		}
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Errorf("Expected %s, got %s", expected, cards[0])
	}
}

func TestSort(t *testing.T) {
	// Cards sorted in reverse
	cards := New(Sort(func(cards []Card) func(i, j int) bool {
		return func(i, j int) bool { return AbsRank(cards[i]) > AbsRank((cards[j])) }
	}))
	expected := Card{Rank: King, Suit: Heart}
	if cards[0] != expected {
		t.Errorf("Expected %s, got %s", expected, cards[0])
	}
}

func TestShuffle(t *testing.T) {
	expectedCards := New(DefaultSort)
	shuffledCards := New(Shuffle)

	if reflect.DeepEqual(shuffledCards, expectedCards) {
		t.Error("Equal value after shuffling")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	jokerCount := 0
	for _, c := range cards {
		if c.Suit == Joker {
			jokerCount++
		}
	}

	if jokerCount != 4 {
		t.Error("Expected 4 Jokers, got", jokerCount)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Four
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Four {
			t.Error("Expected all cards with Rank Two and Four to be filtered out")
		}
	}
}
