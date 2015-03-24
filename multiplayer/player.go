package multiplayer

import "github.com/grzesiek/gonetris/board"

type Player struct {
	Board    *board.Board
	Nickname string
	Host     string
}

func newPlayer(nick, host string, x, y int) *Player {

	player := Player{board.New(x, y), nick, host}
	return &player
}
