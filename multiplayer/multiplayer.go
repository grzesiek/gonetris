package multiplayer

import "sync"

type multiplayer struct {
	Players        []*Player
	NewPlayerEvent chan *Player
	PlayerEvent    chan *Player
}

func New(p int) *multiplayer {

	players := make([]*Player, 0, p)
	newPlayerEvent := make(chan *Player)
	playerEvent := make(chan *Player)

	m := multiplayer{players, newPlayerEvent, playerEvent}
	return &m
}

func (multiplayer *multiplayer) AddPlayer(nick, host string) {
	multiplayer.Players = append(multiplayer.Players, newPlayer(nick, host))
}

func (multiplayer *multiplayer) Handle(wg sync.WaitGroup) {

	defer wg.Done()
}
