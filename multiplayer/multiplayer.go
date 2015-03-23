package main

type multiplayer struct {
	Players        []*Player
	NewPlayerEvent chan *Player
	PlayerEvent    chan *Player
}

func New() *multiplayer {

	players := make([]*Player, 0, Opts.Players)
	newPlayerEvent := make(chan *Player)
	playerEvent := make(chan *Player)

	multiplayer = Multiplayer(players, newPlayerEvent, playerEvent)
	return &multiplayer
}

func (multiplayer *multiplayer) Handle() {

	defer Wg.Done()
	multiplayer.Players = append(Players, newPlayer())
}
