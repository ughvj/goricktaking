package main

import (
	"github.com/ughvj/goricktaking/playingcards"
	"github.com/ughvj/goricktaking/tricktaking"
	"github.com/ughvj/goricktaking/users"
)

func main() {
	deck := playingcards.NewDeck(0)
	players := users.Players{
		*users.NewPlayer("alpha"),
		*users.NewPlayer("bravo"),
		*users.NewPlayer("charlie"),
		*users.NewPlayer("delta"),
	}

	tricktaking.NewPlayBoard(deck.Shuffle(), players)
}
