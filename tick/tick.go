package tick

import (
	"sync"
	"time"
)

type Tick struct {
	Time       time.Duration
	CloseEvent chan bool
}

func New(interval int) *Tick {
	tickTime := time.Duration(interval) * time.Millisecond
	closeEvent := make(chan bool)

	return &Tick{tickTime, closeEvent}
}

func (tick *Tick) Handle(wg sync.WaitGroup, brickOperationEvent chan string) {

	defer wg.Done()

	for {
		select {
		case <-tick.CloseEvent:
			return
		default:
			brickOperationEvent <- "BrickMoveDown"
			time.Sleep(tick.Time)
		}
	}
}
