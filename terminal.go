package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Position struct {
	X int
	Y int
}

func init() {

	termbox.Init()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Sync()
}

func PrintText(text string, p Position) {

	for i, char := range text {
		termbox.SetCell(p.X+i, p.Y, char, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func HandleTerminal() {

	defer Wg.Done()
	defer termbox.Close()

	for Running {

		termbox.Flush()
		time.Sleep(100 * time.Millisecond)
	}
}
