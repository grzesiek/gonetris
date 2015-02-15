package main

import (
	"time"
)

var (
	Tick time.Duration
)

func init() {
	Tick = 200 * time.Millisecond
}

func HandleTick() {

	defer Wg.Done()

	for Running {
		if !Paused {
			BoardEvent <- BrickMoveDown
		}

		time.Sleep(Tick)
	}
}
