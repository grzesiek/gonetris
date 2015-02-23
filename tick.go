package main

import (
	"time"
)

var (
	tick      time.Duration
	TickChan  = make(chan bool)
	TickClose = make(chan bool)
)

func init() {
	tick = 200 * time.Millisecond
}

func HandleTick() {

	defer Wg.Done()

	for {
		select {
		case <-TickClose:
			return
		default:
			TickChan <- true
			time.Sleep(tick)
		}
	}
}
