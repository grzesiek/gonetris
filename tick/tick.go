package tick

import (
	"sync"
	"time"
)

type Tick struct {
	Time       time.Duration
	TickEvent  chan bool
	CloseEvent chan bool
}

func New(interval int) *Tick {
	tickTime := time.Duration(interval) * time.Millisecond
	closeEvent := make(chan bool)
	tickEvent := make(chan bool)

	return &Tick{tickTime, tickEvent, closeEvent}
}

func (tick *Tick) Handle(wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		select {
		case <-tick.CloseEvent:
			return
		default:
			tick.TickEvent <- true
			time.Sleep(tick.Time)
		}
	}
}
