package main

import (
	"time"
)

var (
	tick      time.Duration
	TickClose = make(chan bool)
)

func HandleTick(interval int) {

	defer Wg.Done()

        tick := time.Duration(interval) * time.Millisecond

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
