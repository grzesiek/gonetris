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

	/* First, random brick */
	brick := NextBrick()
	BricksChan <- brick

	for Running {

		if !Paused {
			select {
			case newBrick := <-BricksChan:
				brick = newBrick
			default:
			}
			brick.MoveDown()
			BoardEvent <- MyPlayer.Board
		}
		time.Sleep(GameTick)
	}
}
