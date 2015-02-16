package main

import (
	"time"
)

var (
	Tick     time.Duration
	TickChan = make(chan bool)
)

func init() {
	Tick = 200 * time.Millisecond
}

func HandleTick() {

	defer Wg.Done()

	for Running {
		if !Paused {
			TickChan <- true
		}

		time.Sleep(Tick)
	}
}
