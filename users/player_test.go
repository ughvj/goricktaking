package users_test

import (
	"fmt"
	"testing"

	"github.com/ughvj/goricktaking/users"
)

func Test_DecideOrder(t *testing.T) {
	players := users.Players{
		*users.NewPlayer("alpha"),
		*users.NewPlayer("bravo"),
		*users.NewPlayer("charlie"),
	}

	startPlayer := players.DecideOrder()

	current := *startPlayer
	first := current.GetName()

	current = *current.NextPlayer()
	second := current.GetName()

	current = *current.NextPlayer()
	third := current.GetName()

	current = *current.NextPlayer()
	loop := current.GetName()

	if first != loop {
		t.Error("NG: no loop")
	}

	if first == second {
		t.Error("NG: first == second")
	}

	if second == third {
		t.Error("NG: second == third")
	}

	t.Log(fmt.Sprintf("%s %s %s %s", first, second, third, loop))
}
