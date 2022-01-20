package playingcards

import (
	"math/rand"
	"time"
)

const (
	// Clover is ...
	Clover int = iota
	// Diamond is ...
	Diamond
	// Heart is ...
	Heart
	// Spade is ...
	Spade
)

// DeckNumExceptJoker is ...
const DeckNumExceptJoker int = 52

// CardNumEachMark is ...
const CardNumEachMark int = 13

type Card struct {
	N int
	S int
}

type Cards []Card

func (cs *Cards) PickAt(index int) *Card {
	if len(*cs) >= index {
		return nil
	}
	picked := &(*cs)[index]
	cs.remove(index)

	return picked
}

func (cs *Cards) Find(n int, s int) int {
	for i, c := range *cs {
		if c.N != n {
			continue
		}
		if c.S != s {
			continue
		}
		return i
	}
	return -1
}

func (cs *Cards) remove(target int) *Cards {
	removed := append((*cs)[:target], (*cs)[target+1:]...)
	cs = &removed
	return cs
}

type Deck struct {
	Cards    Cards
	JokerNum int
}

func NewDeck(joker int) Deck {
	var cards Cards
	for i := 0; i < DeckNumExceptJoker; i++ {
		card := Card{
			N: i%CardNumEachMark + 1,
			S: i / CardNumEachMark,
		}
		cards = append(cards, card)
	}
	for j := 0; j < joker; j++ {
		cards = append(cards, Card{
			N: -1,
			S: -1,
		})
	}
	return Deck{
		Cards:    cards,
		JokerNum: joker,
	}
}

func (d *Deck) Shuffle() Deck {
	for i := 0; i < 500; i++ {
		picked := d.pickRandomly()
		d.Cards = append(d.Cards, picked)
	}
	return *d
}

func (d *Deck) Draw(num int) Cards {
	if len(d.Cards) < num {
		return nil
	}
	drawn := Cards{}
	for i := 0; i < num; i++ {
		drawn = append(drawn, d.Cards[0])
		d.remove(0)
	}
	return drawn
}

func (d *Deck) pickRandomly() Card {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(d.Cards))
	picked := d.Cards[index]
	removed := d.remove(index)
	d = removed
	return picked
}

func (d *Deck) remove(target int) *Deck {
	d.Cards = append(d.Cards[:target], d.Cards[target+1:]...)
	return d
}
