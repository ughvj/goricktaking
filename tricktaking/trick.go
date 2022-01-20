package tricktaking

import (
	"github.com/ughvj/goricktaking/playingcards"
	"github.com/ughvj/goricktaking/users"
)

type Trick playingcards.Cards

func (t *Trick) CountNumberOf(n int) (count int) {
	for _, card := range *t {
		if card.N == n {
			count++
		}
	}
	return
}

func (t *Trick) CountSuitOf(s int) (count int) {
	for _, card := range *t {
		if card.S == s {
			count++
		}
	}
	return
}

type Tricks []Trick

func (ts *Tricks) HowMany() int { return len(*ts) }

type PlayBoard struct {
	leaderSuit       int
	Players          users.Players
	PlayerManagement *map[*users.Player]*Management
	firstPlayer      *users.Player
	currentPlayer    *users.Player
}

type Management struct {
	shown      *playingcards.Card
	tricks     Tricks
	hand       playingcards.Cards
	roundPoint int
	totalPoint int
}

func NewPlayBoard(deck playingcards.Deck, players users.Players) *PlayBoard {
	mgmt := make(map[*users.Player]*Management)
	for i := range players {
		mgmt[&players[i]] = &Management{
			hand: deck.Draw(13),
		}
	}

	firstPlayer := players.DecideOrder()

	return &PlayBoard{
		leaderSuit:       -1,
		Players:          players,
		PlayerManagement: &mgmt,
		firstPlayer:      firstPlayer,
		currentPlayer:    firstPlayer,
	}
}

func (pb *PlayBoard) ActorIs() {

}

func NextTurn() {

}
