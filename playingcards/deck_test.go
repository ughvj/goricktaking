package playingcards_test

import (
	"testing"

	"github.com/ughvj/goricktaking/playingcards"
)

func Test_NewDeck(t *testing.T) {
	joker := 2
	d := playingcards.NewDeck(joker)
	for i, card := range d.Cards {
		if i == 52 || i == 53 {
			if card.N != -1 {
				t.Error("NG: Number is different")
			}
			if card.M != -1 {
				t.Error("NG: Mark is different")
			}
			if !card.IsJoker {
				t.Error("NG: IsJoker is false")
			}
			continue
		}
		if card.N != i%playingcards.CardNumEachMark+1 {
			t.Error("NG: Number is different")
		}
		if card.M != i/playingcards.CardNumEachMark {
			t.Error("NG: Mark is different")
		}
		if card.IsJoker {
			t.Error("NG: IsJoker is true")
		}
	}
}
