package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Position struct {
	X int
	Y int
}

var (
	TerminalEvent = make(chan bool)
)

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
	defer fmt.Println("Bye bye !")
	defer termbox.Close()

	for range TerminalEvent {

		termbox.Flush()
	}
}
