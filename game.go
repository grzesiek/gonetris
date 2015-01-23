package main

import (
	"time"
)

var (
	Paused   = false
	GameTick = 400 * time.Millisecond
)

func init() {
}

func HandleGame() {

	defer Wg.Done()

	PrintText("Game started ...", Position{X: 1, Y: 1})

	/* Add first, default player */
	MyPlayer = NewPlayer()
	Players = append(Players, MyPlayer)

	/* Draw first brick */
	NextBrick()

	for Running {

		if !Paused {
			CurrentBrick.MoveDown()
		}
		time.Sleep(GameTick)
	}
}
