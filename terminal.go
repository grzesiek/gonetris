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

func PrintText(value interface{}, p Position) {

	text := fmt.Sprintf("%v", value)
	for i, char := range text {
		termbox.SetCell(p.X+i, p.Y, char, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func HandleTerminal() {

	defer Wg.Done()
	defer fmt.Println("Bye bye !")
	defer termbox.Close()

	for Running {
		select {
		case <-TerminalEvent:
			termbox.Flush()
		}
	}
}
