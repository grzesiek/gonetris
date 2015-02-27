package main

import (
	"github.com/nsf/termbox-go"
	"reflect"
)

var (
	BoardEvent          = make(chan bool)
	BoardBrickOperation = make(chan string)
	BoardClose          = make(chan bool)
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

type BoardBlocker uint16

const (
	BorderLeft BoardBlocker = 1 << iota
	BorderRight
	BorderTop
	BorderBottom
	BrickAtLeft
	BrickAtRight
	BrickBelow
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
	board.BrickNext()

	for {

		select {
		case method := <-BoardBrickOperation:
			/* Player wants to modify brick - move, rotate, drop ... by reflection */
			/* This also handles moving down bick on tick */
			reflect.ValueOf(board).MethodByName(method).Call([]reflect.Value{})
		case <-BoardClose:
			return
		}

		/* Reset empty cells (not filled) */
		board.ResetEmptyCells()
		/* Draw current brick board */
		board.BrickDraw()

		/* User can move birck one last time after it touches something */
		if board.NeedsNextBrick() {
			/* Fill with current brick*/
			board.FillWithBrick()
			/* Chose next brick */
			board.BrickNext()
		}

		/* Draw board */
		TerminalBoardEvent <- *board
	}
}
