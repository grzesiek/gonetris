package main

import (
	"reflect"
)

var ()

type Board struct {
	Matrix     BoardMatrix
	Position   Position
	Shadow     [10]bool
	brick      *Brick
	boardEvent chan Board
	closeEvent chan bool
}

func NewBoard(x, y int) *Board {

	var board Board
	board.Position = Position{X: x, Y: y}
	board.Matrix = NewBoardMatrix()
	board.boardEvent = make(chan Board)
	board.closeEvent = make(chan bool)

	/* TODO: this should go to Brick or be changed in other way
	BoardBrickOperation = make(chan string)
	*/

	/* TODO: This shoudln't be needed anymore
	TerminalNewBoardEvent <- board
	*/

	return &board
}

func (board *Board) Handle(player Player) {

	defer Wg.Done()

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

		/* emit boardEvent */
		boardEvent <- *board
	}
}
