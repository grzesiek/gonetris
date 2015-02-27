package main

import (
	"time"
)

var (
	tick      time.Duration
	TickClose = make(chan bool)
)

func init() {
	tick = 400 * time.Millisecond
}

func HandleTick() {

	defer Wg.Done()

	for {
		select {
		case <-TickClose:
			return
		default:
			BoardBrickOperation <- "BrickMoveDown"
			time.Sleep(tick)
		}
	}
}
