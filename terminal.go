package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func init() {
	termbox.Init()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Sync()
}

func HandleTerminal() {

	defer Wg.Done()
	defer termbox.Close()

	for Running {

		time.Sleep(100 * time.Millisecond)
	}
}
