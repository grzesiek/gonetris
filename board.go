package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"reflect"
	"time"
)

var (
	BoardEvent     = make(chan bool)
	BrickOperation = make(chan string)
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

type BoardBorder uint16

const (
	BorderLeft BoardBorder = iota
	BorderRight
	BorderTop
	BorderBottom
)

func (b Board) Draw() {

	for row, cells := range b.Matrix {
		for col, cell := range cells {
			x, y := b.Position.X+row, b.Position.Y+col
			termbox.SetCell(x, y, cell.Char.Ch, cell.Char.Fg, cell.Char.Bg)
		}
	}
}

func (b Board) DrawFrame() {

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

func (board *Board) BrickTouched(border BoardBorder, move bool) bool {

	brick := board.Brick
	for bx, cells := range brick.Layout {
		for by, cell := range cells {
			x, y := brick.Position.X+(bx*2), brick.Position.Y+by
			if cell == 1 {
				/* Touched right border */
				if border == BorderRight && x+1 == len(board.Matrix)-1 {
					return true
				}
				/* Touched left border */
				if border == BorderLeft && x == 0 {
					return true
				}
				/* Touched bottom border */
				if border == BorderBottom && len(board.Matrix) == y+1 {
					return true
				}
				/* Touched other brick, that already filled board at the bottom */
				if y+1 < len(board.Matrix) && board.Matrix[x][y+1].Filled {
					return true
				}
				if move { /* Check this only if we are moving brick */
					/* Touched other brick, that already filled board at left */
					if x > 2 && board.Matrix[x-2][y].Filled {
						return true
					}
					/* Touched other brick, that already filled board at right */
					if x+2 < len(board.Matrix) && board.Matrix[x+2][y].Filled {
						return true
					}
				}
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

func (board *Board) NextBrick() *Brick {
	rand.Seed(time.Now().UTC().UnixNano())
	brick := &Bricks[rand.Intn(7)]
	brick.Position = Position{0, 0}
	board.Brick = brick
	brick.Board = board

	return brick
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

	TerminalNewBoardEvent <- board

	return &board
}

func HandleBoard() {

	defer Wg.Done()

	/* First player is my player*/
	player := <-PlayersChan
	board := player.Board

	/* Create first brick */
	board.NextBrick()

	for Running {

		select {
		case <-TickChan:
			/* Game tick - move brick down */
			board.Brick.MoveDown()
		case method := <-BrickOperation:
			/* Player wants to modify brick - move, rotate, drop ... by reflection */
			reflect.ValueOf(board.Brick).MethodByName(method).Call([]reflect.Value{})
		}

		/* Reset empty cells (not filled) */
		board.ResetEmptyCells()
		/* Draw current brick board */
		board.DrawBrick()

		/* Draw board */
		TerminalBoardEvent <- *board
	}
}
