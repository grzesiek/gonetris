package main

import (
	"time"
)

var (
	tick      time.Duration
	TickClose = make(chan bool)
)

func HandleTick() {

	defer Wg.Done()

	tick := time.Duration(Opts.Interval) * time.Millisecond

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
