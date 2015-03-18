package main

import ()

type Player struct {
	Board    *Board
	Nickname string
	Host     string
}

var (
	Players     []*Player
	PlayersChan = make(chan *Player)
	PlayerEvent = make(chan *Player)
)

func init() {
	Players = make([]*Player, 0, Opts.Players)
}

func newPlayer() *Player {

	var player Player
	player.Board = NewBoard(5, 5)
	PlayersChan <- &player

	return &player
}

func HandlePlayers() {

	defer Wg.Done()
	Players = append(Players, newPlayer())
}
