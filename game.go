package main

import (
	"time"
)

var (
	Paused   = true
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

		CurrentBrick.MoveDown()
		time.Sleep(GameTick)
	}
}
