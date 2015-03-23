package board

import (
	"reflect"
)

type board struct {
	Matrix     BoardMatrix
	Position   Position
	Shadow     [10]bool
	Brick      *Brick
	BoardEvent chan Board
	CloseEvent chan bool
}

func New(x, y int) *Board {

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

func (board *Board) Handle(wg sync.WaitGroup, player Player) {

	defer wg.Done()

	/* Create first brick */
	board.brickNext()

	for {

		select {
		case method := <-BoardBrickOperation:
			/* Player wants to modify brick - move, rotate, drop ... by reflection */
			/* This also handles moving down bick on tick */
			reflect.ValueOf(board).MethodByName(method).Call([]reflect.Value{})
		case <-Close:
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
