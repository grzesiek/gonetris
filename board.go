package main

import (
	"reflect"
)

var (
	BoardEvent          = make(chan bool)
	BoardBrickOperation = make(chan string)
	BoardClose          = make(chan bool)
)

type BoardCell struct {
	Color    Color
	Empty    bool
	Embedded bool
}

type Board struct {
	Matrix   [10][20]BoardCell
	Position Position
	Brick    *Brick
}

type BrickBlocker uint16

const (
	BorderLeft BrickBlocker = 1 << iota
	BorderRight
	BorderTop
	BorderBottom
	BrickAtLeft
	BrickAtRight
	BrickBelow
	Something = 127
)

func (b *Board) ResetEmptyCells() {

	for x, cells := range b.Matrix {
		for y, cell := range cells {
			if cell.Embedded == false {
				b.Matrix[x][y].Empty = true
				b.Matrix[x][y].Color = ColorBlack
			}
		}
	}
}

func NewBoard(x, y int) *Board {

	var board Board

	board.Position = Position{X: x, Y: y}

	for x, cells := range board.Matrix {
		for y := range cells {
			board.Matrix[x][y].Color = ColorBlack
			board.Matrix[x][y].Empty = true
			board.Matrix[x][y].Embedded = false
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
