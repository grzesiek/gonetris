package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Color termbox.Attribute

const (
	ColorDefault Color = (Color)(termbox.ColorDefault)
	ColorBlack         = (Color)(termbox.ColorBlack)
	ColorRed           = (Color)(termbox.ColorRed)
	ColorGreen         = (Color)(termbox.ColorGreen)
	ColorYellow        = (Color)(termbox.ColorYellow)
	ColorBlue          = (Color)(termbox.ColorBlue)
	ColorMagenta       = (Color)(termbox.ColorMagenta)
	ColorCyan          = (Color)(termbox.ColorCyan)
	ColorWhite         = (Color)(termbox.ColorWhite)
)

type Position struct {
	X int
	Y int
}

var (
	TerminalNewBoardEvent = make(chan Board)
	TerminalBoardEvent    = make(chan Board)
	TerminalClose         = make(chan bool)
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

	for {
		select {
		case board := <-TerminalNewBoardEvent:
			board.DrawFrame()
		case board := <-TerminalBoardEvent:
			board.Draw()
		case <-TerminalClose:
			return
		}

		termbox.Flush()
	}
}
