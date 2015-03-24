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

func (multiplayer *multiplayer) AddPlayer(nick, host string, x, y int) *Player {

	player := newPlayer(nick, host, x, y)
	multiplayer.Players = append(multiplayer.Players, player)

	return player
}

func (multiplayer *multiplayer) Handle(wg sync.WaitGroup) {

	/* multiplayer TODO */
	defer wg.Done()
}
