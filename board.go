package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Board struct {
	Matrix   [20][20]rune
	Position Position
}

func (b *Board) Draw() {

	for row, cells := range b.Matrix {
		for col, _ := range cells {
			x, y := b.Position.X+row, b.Position.Y+col
			termbox.SetCell(x, y, '#', termbox.ColorRed, termbox.ColorGreen)
		}
	}

}

func (b *Board) DrawFrame() {

	width, height := len(b.Matrix), len(b.Matrix[0])
	x, y := b.Position.X, b.Position.Y
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

func NewBoard(x, y int) *Board {

	var board Board

	board.Position = Position{X: x, Y: y}
	board.DrawFrame()

	return &board
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
