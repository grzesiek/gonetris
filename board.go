package main

import (
	"github.com/nsf/termbox-go"
)

var (
	BoardEvent = make(chan *Board)
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

func (b *Board) DrawBrick(brick *Brick) {

	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+(bx*2), brick.Position.Y+by
			if cell == 1 {
				b.Matrix[x][y].Char.Ch = '['
				b.Matrix[x+1][y].Char.Ch = ']'
				b.Matrix[x][y].Char.Bg = brick.Color
				b.Matrix[x+1][y].Char.Bg = brick.Color
				b.Matrix[x][y].Char.Fg = termbox.ColorBlack
				b.Matrix[x+1][y].Char.Fg = termbox.ColorBlack
			}
		}
	}

}

func (board *Board) BrickSticked(brick *Brick) (sticked bool) {

	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+(bx*2), brick.Position.Y+by
			if cell == 1 && len(board.Matrix) == y+1 {
				return true
			}
			if cell == 1 && board.Matrix[x][y+1].Filled {
				return true
			}
		}
	}

	return false
}

func (board *Board) FillWithBrick(brick *Brick) {

	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+(bx*2), brick.Position.Y+by
			if cell == 1 {
				board.Matrix[x][y].Filled = true
				board.Matrix[x+1][y].Filled = true
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

	player := <-PlayerChan
	brick := GetBrick()

	for board := range BoardEvent {
		/* This is MyPlayer's board event */
		if board == player.Board {

			/* Reset empty cells (not filled) */
			board.ResetEmptyCells()

			/* Draw current brick on MyPlayer's board */
			board.DrawBrick(brick)

			/* Check brick position */
			if board.BrickSticked(brick) {
				board.FillWithBrick(brick)
				BrickNew <- true
				brick = GetBrick()
			}
		}

		board.Draw()
		TerminalEvent <- true
	}
}
