package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func HandleBoard() {

	termbox.HideCursor()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Sync()

	termbox.SetCell(0, 0, 'a', termbox.ColorRed, termbox.ColorGreen)
	termbox.SetCell(10, 10, 'b', termbox.ColorRed, termbox.ColorGreen)

	termbox.Flush()

	for Running {
		time.Sleep(10 * time.Millisecond)
	}

}
