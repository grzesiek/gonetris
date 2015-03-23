package multiplayer

type Player struct {
	Nickname string
	Host     string
}

func newPlayer(nick, host string) *Player {

	player := Player{nick, host}
	return &player
}
