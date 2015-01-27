package main

import (
	"time"
)

var (
	GoGame *Game
)

type Game struct {
	Paused bool
	Tick   time.Duration
}

func init() {
	GoGame = &Game{
		Paused: false,
		Tick:   (400 * time.Millisecond)}
}

/* Add first, default player */
func (game Game) AddFirstPlayer() {

	MyPlayer = NewPlayer()
	Players = append(Players, MyPlayer)
}

func (game *Game) Pause() {
	game.Paused = true
}

func (game *Game) Loop() {

	for Running {

		if Running && !game.Paused {
			BrickEventChan <- BrickMoveDown
			BoardEventChan <- MyPlayer.Board
		}

		time.Sleep(game.Tick)
	}
}

func HandleGame() {

	defer Wg.Done()

	PrintText("Game started ...", Position{X: 1, Y: 1})

	GoGame.AddFirstPlayer()
	GoGame.Loop()

}
