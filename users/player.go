package users

import (
	"math/rand"
	"time"
)

var id int = 0

type Player struct {
	id         int
	name       string
	nextPlayer *Player
}

func NewPlayer(name string) *Player {
	p := Player{
		id:         id,
		name:       name,
		nextPlayer: nil,
	}
	id++
	return &p
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) NextPlayer() *Player {
	return p.nextPlayer
}

type Players []Player

func (ps *Players) getIDList() []int {
	l := []int{}
	for _, p := range *ps {
		l = append(l, p.id)
	}
	return l
}

func (ps *Players) findByID(id int) *Player {
	for _, p := range *ps {
		if p.id == id {
			return &p
		}
	}
	return nil
}

func (ps *Players) DecideOrder() (start *Player) {
	rand.Seed(time.Now().UnixNano())
	ids := ps.getIDList()
	index := rand.Intn(len(ids))
	start = ps.findByID(ids[index])
	prev := start
	ids = append(ids[:index], ids[index+1:]...)
	for len(ids) != 0 {
		index = rand.Intn(len(ids))
		prev.nextPlayer = ps.findByID(ids[index])
		prev = prev.nextPlayer
		ids = append(ids[:index], ids[index+1:]...)
	}
	prev.nextPlayer = start
	return
}
