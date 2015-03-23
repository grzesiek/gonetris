package main

import (
	"time"
)

type Tick struct {
	tickTime   time.Duration
	closeEvent chan bool
}

func NewTick(interval int) *Tick {
	tickTime := time.Duration(interval) * time.Millisecond
	closeEvent := make(chan bool)

	return &Tick{tickTime, closeEvent}
}

func (tick *Tick) Handle(brickOperationEvent chan string) {

	defer Wg.Done()

	for {
		select {
		case <-closeEvent:
			return
		default:
			brickOperationEvent <- "BrickMoveDown"
			time.Sleep(time)
		}
	}
}
