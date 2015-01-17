package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

var (
	PlayerBoard = new(Board)
)

type Board struct {
	Runes [22][22]rune
	X     int
	Y     int
}

func (b *Board) Draw() {

	for row, runes := range PlayerBoard.Runes {
		for col, _ := range runes {
			x, y := PlayerBoard.X+row, PlayerBoard.Y+col
			termbox.SetCell(x, y, '#', termbox.ColorRed, termbox.ColorGreen)
		}
	}
	termbox.Flush()

}

func init() {
	PlayerBoard.X = 6
	PlayerBoard.Y = 4
}

func HandleBoard() {

	defer Wg.Done()

	for Running {

		PlayerBoard.Draw()
		time.Sleep(100 * time.Millisecond)
	}

}
