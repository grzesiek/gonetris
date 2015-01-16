package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func init() {
	termbox.Init()
}

func HandleTerminal() {

	defer termbox.Close()
	for Running {

		time.Sleep(100 * time.Millisecond)
	}
	Wg.Done()
}
