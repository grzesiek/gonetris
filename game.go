package main

import (
	"time"
)

var (
	GoGame Game
)

type Game struct {
	Paused bool
	Tick   time.Duration
}

func init() {
	GoGame = Game{
		Paused: false,
		Tick:   (400 * time.Millisecond)}
}

func (game Game) Loop() {

	for Running {

		if !game.Paused {
			game.MoveDownBrick()
			BoardEvent <- MyPlayer.Board
		}

		time.Sleep(game.Tick)
	}
}

/* Add first, default player */
func (game Game) AddFirstPlayer() {

	MyPlayer = NewPlayer()
	Players = append(Players, MyPlayer)
}

func (game Game) GetBrick() *Brick {
	BrickGet <- true
	return <-BricksChan
}

func (game Game) MoveDownBrick() {
	BrickDown <- true
}

func (game Game) NewBrick() {
	BrickNew <- true
}

func HandleGame() {

	defer Wg.Done()

	PrintText("Game started ...", Position{X: 1, Y: 1})

	GoGame.AddFirstPlayer()
	GoGame.Loop()

}
