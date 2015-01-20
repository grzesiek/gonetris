package main

import (
	"time"
)

func init() {
}

func HandleGame() {

	defer Wg.Done()

	// Add first, default player
	Players = append(Players, NewPlayer())

	PrintDebug("debug message")

	for Running {

		time.Sleep(100 * time.Millisecond)
	}
}
