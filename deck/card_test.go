package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Diamond, Rank: Ace})

	//Output:
	//Ace of Diamonds

}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error(fmt.Sprintf("Wrong size of cards in deck. Expected: %d, Actual: %d", 52, len(cards)))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Suit: Spade, Rank: Ace}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades. Got:", cards[0])
	}
}
