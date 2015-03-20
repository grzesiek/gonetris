package main

import (
	"reflect"
)

var (
	BoardEvent          = make(chan bool)
	BoardBrickOperation = make(chan string)
	BoardClose          = make(chan bool)
)

type Board struct {
	Matrix   BoardMatrix
	Position Position
	Brick    *Brick
	Shadow   [10]bool
}

func NewBoard(x, y int) *Board {

	var board Board
	board.Position = Position{X: x, Y: y}
	board.Matrix = NewBoardMatrix()

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
		board.Matrix.ResetEmptyCells()
		/* Draw current brick board */
		board.BrickDraw()
		/* Set current brick shadow */
		board.BrickSetShadow()

		/* User can move birck one last time after it touches something */
		if board.NeedsNextBrick() {

			/* Fill with current brick*/
			board.FillWithBrick()
			/* Chose next brick */
			board.BrickNext()
			/* Remove full lines */
			board.Matrix.RemoveFullLines()
		}

		/* Draw board */
		TerminalBoardEvent <- *board
	}
}
