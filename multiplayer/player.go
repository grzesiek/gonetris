package main

import (
	"github.com/grzesiek/gonetris/board"
)

type Player struct {
	Board    *Board
	Nickname string
	Host     string
}

func newPlayer() *Player {

	var player Player
	player.Board = NewBoard(5, 5)
	PlayersChan <- &player

	return &player
}
