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

func drawBoardFrame(board Board) {

	width, height := len(board.Matrix)*2, len(board.Matrix[0])
	x, y := board.Position.X, board.Position.Y

	for i := -1; i <= width; i++ {
		ch := '-'
		if i == -1 || i == width {
			ch = '+'
		}
		termbox.SetCell(x+i, y-1, ch, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x+i, y+height, ch, termbox.ColorWhite, termbox.ColorBlack)
	}
	for i := 0; i < height; i++ {
		termbox.SetCell(x-1, y+i, '|', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x+width, y+i, '|', termbox.ColorWhite, termbox.ColorBlack)
	}
}

func drawBoard(board Board) {

	for row, cells := range board.Matrix {
		for col, cell := range cells {
			x, y := board.Position.X+(row*2), board.Position.Y+col
			termbox.SetCell(x, y, '[', termbox.ColorBlack, (termbox.Attribute)(cell.Color))
			termbox.SetCell(x+1, y, ']', termbox.ColorBlack, (termbox.Attribute)(cell.Color))
		}
	}

}

func HandleTerminal() {

	defer Wg.Done()
	defer fmt.Println("Bye bye !")
	defer termbox.Close()

	for {
		select {
		case board := <-TerminalNewBoardEvent:
			drawBoardFrame(board)
		case board := <-TerminalBoardEvent:
			drawBoard(board)
		case <-TerminalClose:
			return
		}

		termbox.Flush()
	}
}
