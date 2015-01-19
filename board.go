package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Board struct {
	Runes [22][22]rune
	X     int
	Y     int
}

func (b *Board) Draw() {

	for row, runes := range b.Runes {
		for col, _ := range runes {
			x, y := b.X+row, b.Y+col
			termbox.SetCell(x, y, '#', termbox.ColorRed, termbox.ColorGreen)
		}
	}

}

func HandleBoards() {

	defer Wg.Done()

	for Running {

		for _, player := range Players {
			player.Board.Draw()
		}
		time.Sleep(100 * time.Millisecond)
	}

}
