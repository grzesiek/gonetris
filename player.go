package main

import ()

type Player struct {
	Board    *Board
	Nickname string
	Host     string
}

var (
	Players  []*Player
	MyPlayer *Player
)

func init() {
	Players = make([]*Player, 0, opts.Players)
}

func NewPlayer() *Player {

	var player Player
	player.Board = NewBoard(5, 5)

	return &player
}
