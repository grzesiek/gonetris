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

	for Running {

		if !Paused {
			BrickDown <- true
			BoardEvent <- MyPlayer.Board
		}

		time.Sleep(GameTick)
	}
}
