package main

import (
	"time"
)

var (
	Paused   = true
	GameTick = 1000 * time.Millisecond
)

func init() {
}

func GameNextStep() {
}

func HandleGame() {

	defer Wg.Done()

	PrintText("Game started ...", Position{X: 1, Y: 1})

	/* Add first, default player */
	MyPlayer = NewPlayer()
	Players = append(Players, MyPlayer)

	NextBrick()

	for Running {

		CurrentBrick.MoveDown()

		GameNextStep()
		time.Sleep(GameTick)
	}
}
