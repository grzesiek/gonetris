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

func HandleGame() {

	defer Wg.Done()

	/* Add first, default player */
	Players = append(Players, NewPlayer())

	PrintStatus("Game started ...")

	for Running {

		//		BrickNext()
		time.Sleep(GameTick)
	}

}
