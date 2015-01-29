package main

import (
	"time"
)

func init() {
	Tick = 200 * time.Millisecond
}

func HandleGame() {

	defer Wg.Done()

	PrintText("Game started ...", Position{X: 1, Y: 1})
	Players = append(Players, NewPlayer())

	for Running {

		if !Paused {
			BoardEvent <- BrickMoveDown
		}

		time.Sleep(Tick)
	}
}
