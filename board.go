package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

var (
	BoardEvent = make(chan Event)
)

type BoardCell struct {
	Char   termbox.Cell
	Filled bool
}

type Board struct {
	Matrix   [20][20]BoardCell
	Position Position
	Brick    *Brick
}

type BorderType uint16

const (
	BorderLeft BorderType = iota
	BorderRight
)

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

func (b *Board) DrawBrick() {

	brick := b.Brick
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

func (board *Board) MoveBrickDown() {
	board.Brick.MoveDown()
}

func (board *Board) BrickSticked() (sticked bool) {

	brick := board.Brick
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

func (board *Board) FillWithBrick() {

	brick := board.Brick
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

func (board *Board) BrickTouches(border BorderType) bool {

	brick := board.Brick
	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, _ := brick.Position.X+(bx*2), brick.Position.Y+by
			if cell == 1 {
				if border == BorderRight && x+1 == len(board.Matrix)-1 {
					return true
				}
				if border == BorderLeft && x == 0 {
					return true
				}
			}
		}
	}

	return false
}

func (board *Board) NextBrick() *Brick {
	rand.Seed(time.Now().UTC().UnixNano())
	board.Brick = &Bricks[rand.Intn(7)]
	board.Brick.Position = Position{0, 0}
	return board.Brick
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
	board := player.Board
	board.NextBrick()

	for event := range BoardEvent {

		switch event {
		case BrickMoveDown:
			/* Check if bricked touch something */
			if board.BrickSticked() {
				/* Fill with current brick*/
				board.FillWithBrick()
				/* Chose next brick */
				board.NextBrick()
			} else {
				board.Brick.MoveDown()
			}
		case BrickMoveLeft:
			if !board.BrickTouches(BorderLeft) {
				board.Brick.MoveLeft()
			}
		case BrickMoveRight:
			if !board.BrickTouches(BorderRight) {
				board.Brick.MoveRight()
			}
		}

		/* Reset empty cells (not filled) */
		board.ResetEmptyCells()
		/* Draw current brick on MyPlayer's board */
		board.DrawBrick()

		for _, p := range Players {
			p.Board.Draw()
		}

		TerminalEvent <- true
	}
}
