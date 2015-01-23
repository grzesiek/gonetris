package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type BoardCell struct {
	Char   termbox.Cell
	Filled bool
}

type Board struct {
	Matrix   [20][20]BoardCell
	Position Position
}

func (b *Board) Draw() {

	for row, cells := range b.Matrix {
		for col, cell := range cells {
			x, y := b.Position.X+row, b.Position.Y+col
			termbox.SetCell(x, y, cell.Char.Ch, cell.Char.Fg, cell.Char.Bg)
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

func (b *Board) ResetEmptyCells() {

	for x, cells := range b.Matrix {
		for y, cell := range cells {
			if cell.Filled == false {
				b.Matrix[x][y].Char.Fg = termbox.ColorDefault
				b.Matrix[x][y].Char.Bg = termbox.ColorBlack
				b.Matrix[x][y].Char.Ch = ' '
			}
		}
	}
}

func NewBoard(x, y int) *Board {

	var board Board

	board.Position = Position{X: x, Y: y}

	for x, cells := range board.Matrix {
		for y := range cells {
			board.Matrix[x][y].Char.Fg = termbox.ColorDefault
			board.Matrix[x][y].Char.Bg = termbox.ColorBlack
			board.Matrix[x][y].Filled = false
		}
	}

	board.DrawFrame()

	return &board
}

func HandleBoards() {

	defer Wg.Done()

	for Running {

		//TODO

		if MyPlayer != nil {
			/* Reset empty cells (not filled) */
			MyPlayer.Board.ResetEmptyCells()
		}

		if CurrentBrick != nil {
			/* Draw current brick on MyPlayer's board */
			CurrentBrick.DrawOnBoard()
		}

		/* Draw all player's boards*/
		for _, player := range Players {
			player.Board.Draw()
		}

		time.Sleep(100 * time.Millisecond)
	}

}
