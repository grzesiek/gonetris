package main

type Multiplayer struct {
	Players        []*Player
	NewPlayerEvent chan *Player
	PlayerEvent    chan *Player
}

func NewMultiplayer() *Multiplayer {

	players := make([]*Player, 0, Opts.Players)
	newPlayerEvent := make(chan *Player)
	playerEvent := make(chan *Player)

	multiplayer = Multiplayer(players, newPlayerEvent, playerEvent)
	return &multiplayer
}

func (multiplayer *Multiplayer) Handle() {

	defer Wg.Done()
	Players = append(Players, newPlayer())
}
