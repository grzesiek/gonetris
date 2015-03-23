package board

import (
	"github.com/grzesiek/gonetris/brick"
	"reflect"
)

type position struct {
	X, Y int
}
type board struct {
	Matrix     matrix
	Shadow     [10]bool
	Brick      *brick.Brick
	BoardEvent chan Board
	CloseEvent chan bool
	X          int
	Y          int
}

func New(x, y int) *Board {

	var board Board
	board.X = x
	board.Y = y
	board.Matrix = newMatrix()
	board.BoardEvent = make(chan Board)
	board.CloseEvent = make(chan bool)

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
		board.Matrix.resetEmptyCells()
		/* Draw current brick board */
		board.brickDraw()
		/* Set current brick shadow */
		board.brickSetShadow()

		/* User can move birck one last time after it touches something */
		if board.needsNextBrick() {

			/* Fill with current brick*/
			board.fillWithBrick()
			/* Chose next brick */
			board.brickNext()
			/* Remove full lines */
			board.Matrix.removeFullLines()
		}

		/* emit boardEvent */
		boardEvent <- *board
	}
}
