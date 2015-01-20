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

func PrintDebug(text string) {

	for i, char := range text {
		termbox.SetCell(2+i, 2, char, termbox.ColorWhite, termbox.ColorRed)
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
